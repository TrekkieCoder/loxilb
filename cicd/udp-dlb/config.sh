#!/bin/bash

source ../common.sh

echo "#########################################"
echo "Spawning all hosts"
echo "#########################################"

spawn_docker_host --dock-type loxilb --dock-name llb1 --extra-args "--cluster=172.17.0.4,172.17.0.5"
spawn_docker_host --dock-type loxilb --dock-name llb2 --extra-args "--cluster=172.17.0.3,172.17.0.5"
spawn_docker_host --dock-type loxilb --dock-name llb3 --extra-args "--cluster=172.17.0.3,172.17.0.4"
spawn_docker_host --dock-type host --dock-name l3h1
spawn_docker_host --dock-type host --dock-name l3h2
spawn_docker_host --dock-type host --dock-name l3h3
spawn_docker_host --dock-type host --dock-name l3h4
spawn_docker_host --dock-type host --dock-name l3ep1
spawn_docker_host --dock-type host --dock-name l3ep2
spawn_docker_host --dock-type host --dock-name l3ep3
spawn_docker_host --dock-type host --dock-name br1
spawn_docker_host --dock-type host --dock-name br2
spawn_docker_host --dock-type host --dock-name br3

echo "#########################################"
echo "Connecting and configuring  hosts"
echo "#########################################"

connect_docker_hosts l3h1 br1
connect_docker_hosts l3h2 br1
connect_docker_hosts l3h3 br1
connect_docker_hosts l3h4 br1
connect_docker_hosts br1 llb1
connect_docker_hosts br1 llb2
connect_docker_hosts br1 llb3
connect_docker_hosts l3ep1 br2
connect_docker_hosts l3ep2 br3
connect_docker_hosts llb1 br2
connect_docker_hosts llb1 br3
connect_docker_hosts llb2 br2
connect_docker_hosts llb2 br3
connect_docker_hosts llb3 br2
connect_docker_hosts llb3 br3

#L3 config
config_docker_host --host1 l3h1 --host2 br1 --ptype phy --addr 10.10.10.1/24 --gw 10.10.10.254
config_docker_host --host1 l3h2 --host2 br1 --ptype phy --addr 10.10.10.2/24 --gw 10.10.10.253
config_docker_host --host1 l3h3 --host2 br1 --ptype phy --addr 10.10.10.3/24 --gw 10.10.10.252
config_docker_host --host1 l3h4 --host2 br1 --ptype phy --addr 10.10.10.4/24 --gw 10.10.10.252
config_docker_host --host1 l3ep1 --host2 br2 --ptype phy --addr 31.31.31.1/24 --gw 31.31.31.254
config_docker_host --host1 l3ep2 --host2 br3 --ptype phy --addr 32.32.32.1/24 --gw 32.32.32.254
config_docker_host --host1 llb1 --host2 br1 --ptype phy --addr 10.10.10.254/24
config_docker_host --host1 llb2 --host2 br1 --ptype phy --addr 10.10.10.253/24
config_docker_host --host1 llb3 --host2 br1 --ptype phy --addr 10.10.10.252/24
config_docker_host --host1 llb1 --host2 br2 --ptype phy --addr 31.31.31.254/24
config_docker_host --host1 llb1 --host2 br3 --ptype phy --addr 32.32.32.254/24
config_docker_host --host1 llb2 --host2 br2 --ptype phy --addr 31.31.31.253/24
config_docker_host --host1 llb2 --host2 br3 --ptype phy --addr 32.32.32.253/24
config_docker_host --host1 llb3 --host2 br2 --ptype phy --addr 31.31.31.252/24
config_docker_host --host1 llb3 --host2 br3 --ptype phy --addr 32.32.32.252/24

create_docker_host_cnbridge --host1 br1 --host2 llb1
create_docker_host_cnbridge --host1 br1 --host2 llb2
create_docker_host_cnbridge --host1 br1 --host2 llb3
create_docker_host_cnbridge --host1 br1 --host2 l3h1
create_docker_host_cnbridge --host1 br1 --host2 l3h2
create_docker_host_cnbridge --host1 br1 --host2 l3h3
create_docker_host_cnbridge --host1 br1 --host2 l3h4
create_docker_host_cnbridge --host1 br2 --host2 l3ep1
create_docker_host_cnbridge --host1 br2 --host2 llb1
create_docker_host_cnbridge --host1 br2 --host2 llb2
create_docker_host_cnbridge --host1 br2 --host2 llb3
create_docker_host_cnbridge --host1 br3 --host2 l3ep2
create_docker_host_cnbridge --host1 br3 --host2 llb1
create_docker_host_cnbridge --host1 br3 --host2 llb2
create_docker_host_cnbridge --host1 br3 --host2 llb3

# Create LB rule
#$dexec llb1 loxicmd create lb 20.20.20.1 --udp=2020:8080 --endpoints=31.31.31.1:1,32.32.32.1:1,33.33.33.1:1 --select=persist --mode=onearm
$dexec llb1 loxicmd create lb 20.20.20.1 --tcp=2020:8080-8081 --endpoints=31.31.31.1:1,32.32.32.1:1 --select=persist --mode=onearm
$dexec llb2 loxicmd create lb 30.30.30.1 --tcp=2020:8080-8081 --endpoints=31.31.31.1:1,32.32.32.1:1 --select=persist --mode=onearm
$dexec llb3 loxicmd create lb 40.40.40.1 --tcp=2020:8080-8081 --endpoints=31.31.31.1:1,32.32.32.1:1 --select=persist --mode=onearm
