variable "aws_region" { type = string }
variable "aws_account_number" { type = string }
variable "access_key" { type = string }
variable "secret_key" { type = string }
variable "aws_bucket_name" { type = string }
variable "project_name" { type = string }
variable "importer_source_code" { type = string }
variable "generator_source_code" { type = string }
variable "indexer_source_code" { type = string }
variable "prof_cursoeduca_username" { type = string }
variable "prof_cursoeduca_password" { type = string }
variable "prof_cursoeduca_base_url" { type = string }
variable "cursoeduca_api_key" { type = string }
variable "class_cursoeduca_base_url" { type = string }
variable "ddb_name" { type = string }
variable "lambda_generator_timeout" {
  type    = number
  default = 60
}
variable "lambda_days_log_retention" {
  type    = number
  default = 90
}
variable "lambda_indexer_timeout" {
  type    = number
  default = 30
}