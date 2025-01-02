. ../common.sh
rm -fr iperf-*.log
$hexec l3h1 iperf3 -c 20.20.20.1 -p 2020 -u -t25 --logfile iperf-h1.log --forceflush &
#$hexec l3h1 iperf3 -c 20.20.20.1 -p 2020 -u -t25 --logfile iperf-h1-2.log --forceflush &
$hexec l3h2 iperf3 -c 30.30.30.1 -p 2020 -u -t25 --logfile iperf-h2.log --forceflush &
$hexec l3h3 iperf3 -c 40.40.40.1 -p 2020 -u -t25 --logfile iperf-h3.log --forceflush &
$hexec l3h4 iperf3 -c 40.40.40.1 -p 2020 -u -t25 --logfile iperf-h3.log --forceflush &
echo done
