module "dynamodb_certificates" {
  source     = "./modules/dynamodb"
  name       = var.ddb_name
  aws_region = var.aws_region
}

module "sqs_main" {
  source = "./modules/sqs"
  name = var.project_name
  delay_seconds = 30
}

module "lambda_importer" {
  source           = "./modules/lambda"
  name             = "${var.project_name}-importer"
  handler_path     = "bootstrap"
  runtime          = "provided.al2023"
  role_arn         = aws_iam_role.lambda_importer_role.arn
  filename         = var.importer_source_code_hash
  source_code_hash = base64sha256(var.importer_source_code_hash)
  timeout          = 30
  memory_size      = 1024
  log_retention    = 90
  depends_on = [  module.sqs_main ]
  environment = {
    PROF_CURSEDUCA_USERNAME  = var.prof_cursoeduca_username
    PROF_CURSEDUCA_PASSWORD  = var.prof_cursoeduca_password
    PROF_CURSEDUCA_BASE_URL  = var.prof_cursoeduca_base_url
    CURSEDUCA_API_KEY        = var.cursoeduca_api_key
    CLASS_CURSEDUCA_BASE_URL = var.class_cursoeduca_base_url
    AWS_GENERATOR_QUEUE_URL = module.sqs_main.main_url
  }
}