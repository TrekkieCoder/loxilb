#!/bin/bash

source ../common.sh

delete_docker_host llb1
delete_docker_host llb2
delete_docker_host llb3
delete_docker_host br3
delete_docker_host br2
delete_docker_host br1
delete_docker_host l3h1
delete_docker_host l3h2
delete_docker_host l3h3
delete_docker_host l3h4
delete_docker_host l3ep1
delete_docker_host l3ep2
delete_docker_host l3ep3

echo "#########################################"
echo "Deleted testbed"
echo "#########################################"
