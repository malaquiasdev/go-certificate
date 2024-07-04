variable "aws_region" { type = string }
variable "aws_account_number" { type = string }
variable "access_key" { type = string }
variable "secret_key" { type = string }
variable "aws_bucket_name" { type = string }
variable "project_name" { type = string }
variable "importer_source_code" { type = string }
variable "generator_source_code" { type = string }
variable "indexer_source_code" { type = string }
variable "api_get_certificates_source_code" { type = string }
variable "api_get_certificate_file_source_code" { type = string }
variable "prof_cursoeduca_username" { type = string }
variable "prof_cursoeduca_password" { type = string }
variable "prof_cursoeduca_base_url" { type = string }
variable "cursoeduca_api_key" { type = string }
variable "class_cursoeduca_base_url" { type = string }
variable "ddb_name" { type = string }
variable "lambda_generator_timeout" {
  type    = number
  default = 300
}
variable "lambda_days_log_retention" {
  type    = number
  default = 14
}
variable "lambda_indexer_timeout" {
  type    = number
  default = 60
}
variable "cursoeduca_block_list" { 
  type = string 
  default = “423,437,438,439,440,441,444,414,413,408,403,361,360,305,304,303,302,301,300,265,263,262,261,260,251,250,243,242,241,240,239,238,237,236,232,208,207,206,205,204,203,202,201,200,199,198,197,196,195,194,193,192,191,190,189,342,352,346,233,343”
}
}
variable "certificate_url_prefix" { type = string }
