# minikube

```bash
minikube start --kubernetes-version=v1.23.3 --force --image-mirror-country='cn'
minikube dashboard

# https://camel.apache.org/camel-k/1.10.x/installation/installation.html
brew install kamel
kamel install --cluster-setup


# hello 
# https://toscode.gitee.com/apache/camel-k

kamel run hello.groovy
# return: Integration "hello" created
```