variable qwe {
  default = "qwe"
}

variable "some_number" {
  default     = 666
  type        = number
  validation {
    condition = var.some_number == 3
    error_message = "Invalid number (expected 3)."
  }
}