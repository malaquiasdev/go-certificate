variable "aws_region" { type = string }
variable "name" { type = string }
variable "point_in_time_recovery_enabled" {
  description = "Enable point-in-time recovery (backup)"
  type        = bool
  default     = false
}