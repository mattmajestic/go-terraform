// This is a comment
provider "random" {
  version = "~> 3.0"
}

resource "random_pet" "example" {
  length = 4
}

output "pet_name" {
  value = random_pet.example.id
}
