import json
import pprint as pp
import sys
import subprocess as sp

json_data = json.loads("".join(sys.stdin.readlines()))
json_inputs = json.loads(json_data['inputs'])['payload']
json_outputs = json.loads(json_data['outputs'])['payload']
test = json_inputs[0].encode('ascii', 'ignore')

for index, move in enumerate(json_inputs):
    test = move.encode('ascii', 'ignore')
    print test
    print "Actual Result: ",json_outputs[index]
    p = sp.Popen(['go', 'run' ,'mancala2_opt.go'], stdout=sp.PIPE, stdin=sp.PIPE, stderr=sp.PIPE)
    print(p.communicate(input = test)[0])
