## Main Table

| Entity      | Partition Key: PK | Sort Key: SK   |   |   |
|-------------|-------------------|----------------|---|---|
| Certificate | CERTIFICATE#id    | CERTIFICATE#id |   |   |
|             |                   |                |   |   |

### Sample Data

![](https://github.com/malaquiasdev/ekoa-certificate-generator/blob/main/doc/dynamodb/sample/certificate.png?raw=true)

## Global Secondary Index named - GSI1

| Entity      | Partition Key: GSI1PK | Sort Key: GSI1SK   |   |   |
|-------------|-------------------|----------------|---|---|
| Certificate | REPORT#id    | CERTIFICATE#id |   |   |
|             |                   |                |   |   |

### Sample Data

![](https://github.com/malaquiasdev/ekoa-certificate-generator/blob/main/doc/dynamodb/sample/GSI_certificate_GS1PK.png?raw=true)

## Global Secondary Index named - GSI2

| Entity      | Partition Key: GSI2PK | Sort Key: GSI2K   |   |   |
|-------------|-------------------|----------------|---|---|
| Certificate | EMAIL#id    | CERTIFICATE#id |   |   |
|             |                   |                |   |   |

### Sample Data

![](https://github.com/malaquiasdev/ekoa-certificate-generator/blob/main/doc/dynamodb/sample/GSI_certificate_GS2PK.png?raw=true)

## Model NoSQL Workbench Export

![](https://github.com/malaquiasdev/ekoa-certificate-generator/blob/main/doc/dynamodb/model.png?raw=true)

[mode.json](https://github.com/malaquiasdev/ekoa-certificate-generator/blob/main/doc/dynamodb/model.json)