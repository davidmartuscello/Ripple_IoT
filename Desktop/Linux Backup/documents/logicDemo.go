package main

import "fmt"
import "time"

type logEntry struct{
    userid  int
    time    time.Time
    flowRate    float64
    diffConsumed    float64 //The water consumed since the last log was taken
}

type measurePoint struct {
    name string
    room int 
    accessLog []logEntry 
}

type controlPoint struct {
    name    string
    room    int 
    state   bool
    //accessLog []logEntry
}

type user struct {
    name    string
    userid  int
    admin   bool
    waterUsed   float64
}

//-------------------------------------------
    
func (mp *measurePoint) newLogEntry(logUser int, fRate float64) float64 {
    
    newLog := logEntry{
        userid: logUser,
        time: time.Now(),
        flowRate: fRate,
    }
    
    mp.accessLog = append(mp.accessLog, newLog)
    
    flowConsumed := mp.measureFlow()
    
    return flowConsumed
}

func (u *user) calcWaterUsed (mpList [1]measurePoint) float64 {
    
    listLen := len(mpList)
    
    for j := 0; j <listLen; j++ {
        mp := mpList[j]     // For each mp in mpList
        
        logLen := len(mp.accessLog)
        for j := 0; j <logLen; j++ {
            logEntry := mp.accessLog[j]     //For each logEntry in the accessLog of that measurePoint
            
            if logEntry.userid == u.userid {    //If the logEntry was made by this user ("u")
                
                u.waterUsed += logEntry.diffConsumed    //add the diffConsumed to the users total "waterUsed" 
            
            }
        }
    }
    return u.waterUsed
}

func (mp *measurePoint) measureFlow () float64 { //calculates diffConsumed and puts it into the mpLog

    logLen := len(mp.accessLog)
    fmt.Printf("logLen:") 
    fmt.Println(logLen)
    if logLen >= 2 {
        lastFlowRate := mp.accessLog[logLen-2].flowRate //What was the flow rate of the last recorded log entry
        
        timeElapsedInt := mp.accessLog[logLen-1].time.Second() - mp.accessLog[logLen-2].time.Second()
        timeElapsed := float64(timeElapsedInt)
        flowConsumed := timeElapsed*lastFlowRate
        mp.accessLog[logLen-1].diffConsumed = flowConsumed
        
        return flowConsumed
        
    } else {
        
        return 0
        //mp.accessLog[logLen].diffConsumed = 0
    }
}
    
func (cp *controlPoint) changeState(isOn bool) bool {
    
    cp.state = isOn
    return cp.state
}

//-------------------------------------------

func main() {
        
        //Define a user
        david := user {
            name:   "David Martuscello",
            userid: 1,
            admin:  true,
            waterUsed:  0,
        }
        
        //Define a valve (controlPoint)
        Valve1 := controlPoint {
            name: "Kitchen Sink Control Valve",
            room: 1,
            state: false,
        }
        
        //Define a flow meter (measurePoint)
        var log []logEntry
        flowMeter1 := measurePoint{
            name:   "Kitchen Sink Flow Meter",
            room:   1,
            accessLog:  log,
        }
        

        //Toggle the state of the controlPoint ON
        fmt.Println(Valve1.changeState(true))
        //create a log in the measurePoint
        fmt.Println(flowMeter1.newLogEntry(1, 2.34))    //Kitchen Sink measuring 2.34 gpm
        
	
        Then := time.Now().Second()
        for time.Now().Second() < (Then+3){     //Homemade 3 second delay - so that water can be consumed as time passes
            //do nothing
        }
        
        //Toggle the state of the controlPoint OFF
        fmt.Println(Valve1.changeState(false))
        //create a log in the measurePoint
        fmt.Println(flowMeter1.newLogEntry(1, 0))       //Return how much water was consumed since last time
        
        
        mpList := [1]measurePoint{flowMeter1}
        //Diplay how much water the user consumed
        fmt.Println(david.calcWaterUsed(mpList))
    
}