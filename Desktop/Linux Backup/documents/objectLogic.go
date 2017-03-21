package object

import "fmt"
import "time"

/*
//Objects-----

//Example Quotas
    //Quota1
        //David can use 600 gallons of showers this month
        //user or group, usageLimit, waterDevice

    //Quota2
        //The family can use 10000 gallons this month
        //user or group, usageLimit
    */
//
/*
NOTES:
 - Should the flow be controllable only by the electronic valve (if you do not have access to a device how to control the water)
 - The user needs to always be able to user their water (i.e. when they dont have their smart phone around, or set their quota too low (or even emergencies where they need to use water))
 - 


*/
//------------------------------------------Objects

type quota struct {     
    name string
    usageLimit int
    waterDevice string // optional 
}

type logEntry struct{
    userid  int
    time    time.Time
    flowRate    float64
    diffConsumed    float64 //The water consumed since the last log was taken (water consumed during this use)
}

type waterDevice struct {
    name    string
    room    int 
    state   bool
    accessLog []logEntry
}

type user struct {
    name    string
    userid  int
    admin   bool
    totalUsage   float64
}

//*************************************************
//------------------------------------------Methods
//*************************************************

// >> waterDevice Methods 
    
/*    
interface waterDevice {
    updateState // When flow reaches a setpoint, change state
    getState    // return true or false for on or off
    getWaterUsed  //amount consumed in this period (since the last state change)
    getFlowRate   //real time flow rate measurement
    getAccessLog  //Returns an array of maps, each map is a logEntry (i.e. {user: "David", time: "1:00, 1,5,2016", consumed: 7})
    //The waterDevice tracks short term usage history...
}
*/

func (wD *waterDevice) newLogEntry(logUser int, fRate float64) float64 {
    
    newLog := logEntry{
        userid: logUser,
        time: time.Now(),
        flowRate: fRate,
    }
    
    wD.accessLog = append(wD.accessLog, newLog)
    
    flowConsumed := wD.measureFlow()
    
    return flowConsumed
}


func (wD *waterDevice) measureFlow () float64 { //calculates diffConsumed and puts it into the mpLog

    logLen := len(wD.accessLog)
    fmt.Printf("logLen:") 
    fmt.Println(logLen)
    if logLen >= 2 {
        lastFlowRate := wD.accessLog[logLen-2].flowRate //What was the flow rate of the last recorded log entry
        
        timeElapsedInt := wD.accessLog[logLen-1].time.Second() - wD.accessLog[logLen-2].time.Second()
        timeElapsed := float64(timeElapsedInt)
        flowConsumed := timeElapsed*lastFlowRate
        wD.accessLog[logLen-1].diffConsumed = flowConsumed
        
        return flowConsumed
        
    } else {
        
        return 0
        //wD.accessLog[logLen].diffConsumed = 0
    }
}


func (wD *waterDevice) updateState( "*FlowData" ) {
        setpoint = 0.4 int //gpm
        hysteresis = 0.2 int //gpm
        if  ((wD.getState() == "off") && (flow > setpoint)){
            wD.setState("on")
        } else if  ((wD.getState() == "on") && (flow < (setpoint-hysteresis))){
            wD.setState("off")
        }

    waterConsumed = calcWaterConsumedJustNow("*FlowData" )
    wD.newLogEntry(1, waterConsumed) //need to change this methods definition above

    //Need a separate "in use" function to consider turning on and off within small amount of time as 1 "use" to be recorded in the log
}

// >> user Methods

 /*
interface user {
    david.useDevice(sink) //change the device state
    david.updateTotalUsage //look in the access log for all devices and save as totalUsage (maybe unneccessary depending on DB implementation)
    david.getTotalUsage // return totalUsage
    david.getDeviceUsage(sink) // look in the sink's access log for all entrys by this user
*/
 
func (u *user) useDevice(sink) 

 
func (u *user) updateTotalUsage (deviceList [1]waterDevice) {
    //Check all devices to add up all the water used by this user
        //NOTE: may eventually implement to look in the database
    
    listLen := len(deviceList)
    
    for j := 0; j <listLen; j++ {
        wD := deviceList[j]     // For each wD in deviceList
        
        logLen := len(wD.accessLog)
        for j := 0; j <logLen; j++ {
            logEntry := wD.accessLog[j]     //For each logEntry in the accessLog of that waterDevice
            
            if logEntry.userid == u.userid {    //If the logEntry was made by this user ("u")
                
                u.totalUsage += logEntry.diffConsumed    //add the diffConsumed to the users total "waterUsed" 
            
            }
        }
    }
}

func (u *user) getTotalUsage (

// >> Quota Methods    
/*
interface quota{
    createQ() // initialize with name, limit and specific device(optional)
    update() // update the quota with the totalConsumed by the user or group
    quotaFilled() // return true if quota is full
}

func (*userQuota u) quotaFilled ( int totalConsumed) quotaFilled bool {
    
    if (quota reached) {
        return quotaFilled = true
    {
}
    
func (*userQuota u) updateQ ( int totalConsumed) quotaFilled bool {
 
    u.waterDevice = device
    u.usageLimit 
    
    if (quota reached) {
        return quotaFilled = true
    {
}
    
*/
//

//------------------------------------------Functions

func initQuota(initName string, initUsageLimit int, initWaterDevice string) quota {
        return quota {
            name string
            usageLimit int
            waterDevice string  
        }
}

func initUser(initName string, initUserID int, initAdmin bool, initWaterUsed int) user {
        return user {
            name:   initName,
            userid: initUserID,
            admin:  initAdmin,
            waterUsed:  initWaterUsed,
        }
}
        
func initWaterDevice(initName string, initRoom int, initState bool) waterDevice {
        var log []logEntry

        return waterDevice {
            name: initName,
            room: initRoom,
            state: initState,
        }
}


    
  
  

//------------------------------------------Main
        
func main() {
    
        //Toggle the state of the controlPoint ON
        fmt.Println(valve1.changeState())
        //create a log in the measurePoint
        fmt.Println(flowMeter1.newLogEntry(1, 2.34))    //Kitchen Sink measuring 2.34 gpm
        
        Then := time.Now().Second()
        for time.Now().Second() < (Then+3){     //Homemade 3 second delay
            //do nothing
        }
        
        //Toggle the state of the controlPoint OFF
        fmt.Println(valve1.changeState())
        //create a log in the measurePoint
        fmt.Println(flowMeter1.newLogEntry(1, 0))       //Return how much water was consumed since last time
        
        
        mpList := [1]measurePoint{flowMeter1}
        //Diplay how much water the user consumed
        fmt.Println(david.calcWaterUsed(mpList))
}
