---
title: Key-Pair
weight: 4603
---

# Multi-Aggregate
This activity is a data-preparation activity for the Multi-Aggregate activity that allows to configure the input object the Multi Aggregate activity needs to work.

## Installation
### Flogo Web
```bash
https://github.com/alexandrev/flogo-components/activity/keypair
```
### Flogo CLI
```bash
flogo install github.com/alexandrev/flogo-components/activity/keypair
```

## Schema
Inputs and Outputs:

```json
{
 "input":[
    {
      "name": "values",
      "type": "complex_object",
      "value":{      
        "metadata":"",              
        "value": "[0.3,0.4,0.5,0.6]"
        }
    },
    {
      "name": "keys",
      "type": "complex_object",
      "value":{      
        "metadata":"",              
        "value": "[\"a\",\"a\",\"a\"]"
        }
    }
  ],
  "output": [
    {
      "name": "values",
      "type": "complex_object",
      "value":{      
        "metadata":"",              
        "value": "[{ \"operation\":\"count\",\"value\": 0.123}]"
        }
      }    
  ]
}
```

## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
|             |          |             |


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