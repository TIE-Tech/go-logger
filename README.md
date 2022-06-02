# go-logger

Public Log Package

# Example of use
```
import  "github.com/TIE-Tech/go-logger"

logger.SetLogger(`{"Console": {"level": "DEBG"}`)

logger.Trace("this is Trace")
logger.Debug("this is Debug")
logger.Info("this is Info")
logger.Warn("this is Warn")
logger.Error("this is Error")
logger.Crit("this is Critical")
logger.Alert("this is Alert")
logger.Emer("this is Emergency")
```

# Log Level

There are currently 8 log output levels, ranging from 0 to 7 from high to bottom. When a certain output level is configured, only logs with a level greater than or equal to this level will be output. Different output adapters support different log level configurations:

|   Level   |   Name   |   Desc   |
| ---- | ---- | ---- |
| 0 | EMER | System level emergency |
| 1 | ALRT | system level warning |
| 2 | CRIT | system level hazard |
| 3 | EROR | User level error |
| 4 | WARN | User level warning |
| 5 | INFO | User level important |
| 6 | DEBG | User level debugging |
| 7 | TRAC | User level basic output |


# Configuration instructions

```
logger.SetLogger("log.json")
logger.SetLogger(`{"Console": {"level": "DEBG"}}`)
```

```
{
    "TimeFormat":"2006-01-02 15:04:05",
    "Console": {
        "level": "TRAC",
        "color": true 
    },
    "File": {
        "filename": "app.log",
        "level": "TRAC",
        "daily": true,
        "maxlines": 1000000,
        "maxsize": 1,
        "maxdays": -1,
        "append": true,
        "permit": "0660"
    },
   "Elastic": {
        "open": false,
        "addr": "http://127.0.0.1:9200",
        "index": "endorse",
        "level": "DEBG"
   }
}
```