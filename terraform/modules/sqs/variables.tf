variable "name" { type = string }
variable "delay_seconds" { type = number }
variable "max_redrive_count" {
  type    = number
  default = 5
}
variable "max_message_size" {
  type    = number
  default = 262144
}
variable "visibility_timeout_seconds" { type = number }