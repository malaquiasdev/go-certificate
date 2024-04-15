## Main Table

| Entity      | Partition Key: PK | Sort Key: SK   |   |   |
|-------------|-------------------|----------------|---|---|
| Certificate | CERTIFICATE#id    | CERTIFICATE#id |   |   |
|             |                   |                |   |   |

### Sample Data

![](doc/dynamodb/sample/certificate.png)

## Global Secondary Index named - GSI1

| Entity      | Partition Key: GSI1PK | Sort Key: GSI1SK   |   |   |
|-------------|-------------------|----------------|---|---|
| Certificate | REPORT#id    | CERTIFICATE#id |   |   |
|             |                   |                |   |   |

### Sample Data

![](doc/dynamodb/sample/GSI_certificate_GS1PK.png)

## Global Secondary Index named - GSI2

| Entity      | Partition Key: GSI2PK | Sort Key: GSI2K   |   |   |
|-------------|-------------------|----------------|---|---|
| Certificate | EMAIL#id    | CERTIFICATE#id |   |   |
|             |                   |                |   |   |

### Sample Data

![](doc/dynamodb/sample/GSI_certificate_GS2PK.png)

## Model NoSQL Workbench Export

![](doc/dynamodb/model.png)

[mode.json](doc/dynamodb/model.json)