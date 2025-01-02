#!/bin/bash
source ../common.sh

rm -fr port*.log
echo SCENARIO-udplb-all-active
for ((i=1;i<3;i++))
do
for ((port=8080;port<=8081;port++))
do
$hexec l3ep$i iperf3 -s -p $port --logfile port$port-ep$i.log 2>&1 >> /dev/null &
done
done

exit
echo "Waiting for servers to start..."
sleep 20 

NUM_HOSTS=3
rm -fr iperf-*.log
$hexec l3h1 iperf3 -c 20.20.20.1 -p 2020 -u -t20 --logfile iperf-h1.log --forceflush &
$hexec l3h1 iperf3 -c 20.20.20.1 -p 2020 -u -t20 --logfile iperf-h1-2.log --forceflush &
$hexec l3h2 iperf3 -c 30.30.30.1 -p 2020 -u -t20 --logfile iperf-h2.log --forceflush &
$hexec l3h3 iperf3 -c 40.40.40.1 -p 2020 -u -t20 --logfile iperf-h3.log --forceflush &

echo "Waiting for tests to finish..."
sleep 60 
code=0
for file in iperf*.log; do
  if grep -q "connected" "$file"; then
    echo "Pass:'connected' found in $file."
  else
    echo "Fail: 'connected' not found in $file."
    code=1
  fi
done
if [[ $code != 0 ]]; then
    echo "SCENARIO-udplb-all-active [FAILED]"
else
    echo "SCENARIO-udplb-all-active [OK]"
fi

sudo pkill -9 iperf3 2>&1 > /dev/null
#rm -fr iperf-*.log
rm -fr port*.log
exit $code
