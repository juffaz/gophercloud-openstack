package main

import (
"fmt"
"github.com/rackspace/gophercloud"
"github.com/rackspace/gophercloud/openstack"
"github.com/rackspace/gophercloud/pagination"
"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
"github.com/rackspace/gophercloud/openstack/compute/v2/images"
"github.com/rackspace/gophercloud/openstack/blockstorage/v2/volumes"
//"github.com/rackspace/gophercloud/openstack/blockstorage/v1/volumetypes"
//"github.com/rackspace/gophercloud/openstack/blockstorage/v2/snapshots"
)

func main() {

opts := gophercloud.AuthOptions{
IdentityEndpoint: "http://control-node:5000/v3",
Username: "usernamez",
DomainName: "default",
Password: "userpassz",
}
provider, err := openstack.AuthenticatedClient(opts)

if err != nil {
fmt.Println("AuthenticatedClient")
fmt.Println(err)
return
}

client, err :=
openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
Region: "DeltaOne",
})
if err != nil {
fmt.Println("NewComputeV2 Error:")
fmt.Println(err)
return
}

clientblockstorage, err :=
openstack.NewBlockStorageV2(provider, gophercloud.EndpointOpts{
Region: "DeltaOne",
})
if err != nil {
fmt.Println("NewBlockStorageV2 Error:")
fmt.Println(err)
return
}



opts2 := servers.ListOpts{}
opts3 := images.ListOpts{}
opts4 := volumes.ListOpts{}
opts5 := snapshots.ListOpts{}
//opts6 := snapshots.CreateOpts{Name: "test1snapikNov1", VolumeID: "dfad5ef0-3786-4c29-8469-cc96916efa36"}
//opts6 := volumetypes.ListOpts{}
//opts6 := snapshots.CreateOpts{Name: "2014_oct", VolumeID: "dfad5ef0-3786-4c29-8469-cc96916efa36"}

//snap, err := snapshots.Create(clientblockstorage, opts6).Extract()


pager := servers.List(client, opts2)
pager2 := images.ListDetail(client, opts3)
pager3 := volumes.List(clientblockstorage, opts4)
pager4 := snapshots.List(clientblockstorage, opts5)
//pager4 := snapshots.Create(clientblockstorage, opts6)
//pager5 := volumetypes.List(clientblockstorage, opts6)



pager.EachPage(func(page pagination.Page) (bool, error) {
serverList, err := servers.ExtractServers(page)
if err != nil {
return false, err
}
for _, s := range serverList {
fmt.Println(s.ID, s.Name, s.Status)
// servers.Delete(client, s.ID);
}
fmt.Println("###  Epta server list  ")
return true, nil
})

pager2.EachPage(func(page pagination.Page) (bool, error) {
imageList, err := images.ExtractImages(page)

if err != nil {
return false, err
}

for _, i := range imageList {

//fmt.Println(i.ID, i.Name, .Status)
fmt.Println(i)
// servers.Delete(client, s.ID);

}
fmt.Println("### Epta image list ")

return true, nil
})


pager3.EachPage(func(page pagination.Page) (bool, error) {
vList, err := volumes.ExtractVolumes(page)

if err !=nil {
return false, err
}

for _, v := range vList {
fmt.Println(v)

}

fmt.Println("###Epta volume list")
return true, nil
})



pager4.EachPage(func(page pagination.Page) (bool, error) {
snapList, err := snapshots.ExtractSnapshots(page)

if err !=nil {
return false, err
}

for _, snap := range snapList {

fmt.Println(snap)

}
fmt.Println("### snap list ###")

return true, nil
})


//snapnew, err := snapshots.Create{clientblockstorage, opts6}.Extract()
//if err != nil {
//		return &snapshots.Snapshot{}, err
//	}
	// There's little value in rewrapping these gophercloud types into yet another abstraction/type, instead just
	// return the gophercloud item
//return snap, nil

//err := snapshots.Create(clientblockstorage, opts6)



//pager5.EachPage(func(page pagination.Page) (bool, error) {
//vtList, err := volumestypes.ExtractVolumeTypes(page)

//if err !=nil {
//return false, err
//}

//for _, vt := range vtList {

//fmt.Println(vt)

//}
//fmt.Println("### volumetypes list ###")

//return true, nil

//})



}
