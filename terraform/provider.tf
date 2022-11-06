terraform {
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = ">= 2.0.0"
    }
  }
}

provider "kubernetes" {
  config_path = "~/.kube/config"
}

resource "kubernetes_namespace" "terraform" {
  metadata {
    name = "terraform"
  }
}

resource "kubernetes_deployment" "terraform" {
  metadata {
    name      = "terraform"
    namespace = kubernetes_namespace.terraform.metadata.0.name
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "MyTerraformApp"
      }
    }
    template {
      metadata {
        labels = {
          app = "MyTerraformApp"
        }
      }
      spec {
        container {
          image = "nginx"
          name  = "nginx-container"
          port {
            container_port = 80
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "terraform" {
  metadata {
    name      = "terraform"
    namespace = kubernetes_namespace.terraform.metadata.0.name
  }
  spec {
    selector = {
      app = kubernetes_deployment.terraform.spec.0.template.0.metadata.0.labels.app
    }
    type = "NodePort"
    port {
      node_port   = 30201
      port        = 80
      target_port = 80
    }
  }
}