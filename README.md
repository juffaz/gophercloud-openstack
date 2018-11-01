# gophercloud-openstack  

for Arch install  
$ yay -S golang  

for Centos 7 install  
$ sudo yum --enablerepo=epel install  golang -y  


$ go get github.com/rackspace/gophercloud  
$ git clone https://github.com/juffaz/gophercloud-openstack.git   
$ cd gopher*   
$ vi openstack-connect.go   


$ go run openstack-connect.go   
or   
$ go build openstack-connect.go   
$ ./openstack-connect  

