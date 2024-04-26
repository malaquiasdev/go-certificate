resource "aws_dynamodb_table" "this" {
  count = var.aws_region == "sa-east-1" ? 1 : 0
  name  = var.name

  hash_key     = "PK"
  billing_mode = "PAY_PER_REQUEST"

  point_in_time_recovery {
    enabled = var.point_in_time_recovery_enabled
  }

  attribute {
    name = "PK"
    type = "S"
  }

  attribute {
    name = "studentEmail"
    type = "S"
  }

  global_secondary_index {
    name            = "studentEmail"
    hash_key        = "studentEmail"
    projection_type = "ALL"
  }
}