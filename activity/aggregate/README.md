---
title: Multi-Aggregate
weight: 4603
---

# Multi-Aggregate
This activity allows you to aggregate data (supporting multi-value aggregation) on a user-define window. A key can be used and you can choose the kind of aggregation operation (count, min, max, avg) you want for each value.

## Installation
### Flogo Web
```bash
https://github.com/alexandrev/flogo-components/activity/aggregate
```
### Flogo CLI
```bash
flogo install github.com/alexandrev/flogo-components/activity/aggregate
```

## Schema
Inputs and Outputs:

```json
{
  "input":[
     {
      "name": "function",
      "type": "string",
      "required": true,
      "allowed" : ["block", "moving", "timeblock"]
    },
    {
      "name": "windowSize",
      "type": "integer",
      "required": true
    },
    {
      "name": "value",
      "type": "complex_object",
      "value":{      
        "metadata":"",              
        "value": "[{ \"operation\":\"count\",\"value\": 0.123}]"
        }
    },
    {
      "name": "key",
      "type": "string"
    }
  ],
  "output": [
    {
      "name": "result",
      "type": "complex_object",
      "value":{      
        "metadata":"",      
        "value": "[2,3,4,5,6]"      
        }
    },
    {
      "name": "report",
      "type": "boolean"
    }
  ]
}
```

## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| function    | True     | The aggregate fuction, currently only average is supported |
| windowSize  | True     | The window size of the values to aggregate |
| value       | False    | The value to aggregate |


## Example
The below example aggregates a 'temperature' attribute with a moving window of size 5:

```json
"id": "aggregate_4",
"name": "Aggregate",
"description": "Simple Aggregator Activity",
"activity": {
  "ref": "github.com/TIBCOSoftware/flogo-contrib/activity/aggregate",
  "input": {
    "function": "average",
    "windowSize": "5"
  },
  "mappings": {
    "input": [
      {
        "type": "assign",
        "value": "temperature",
        "mapTo": "value"
      }
    ]
  }
```