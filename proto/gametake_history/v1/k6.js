import grpc from 'k6/net/grpc';
import { check, sleep } from 'k6';

const client = new grpc.Client();
client.load(['.'], 'gametakehistory.proto');

export const options = {
  hosts: { 'test.k6.io': '1.2.3.4' },
  stages: [
    { duration: '10s', target: 1000 },
    { duration: '10s', target: 2000 },
    { duration: '5', target: 500 },
  ],
  thresholds: { grpc_req_duration: ['avg<100', 'p(95)<50'] },
  noConnectionReuse: true,
  userAgent: 'MyK6UserAgentString/1.0',

};

export default () => {
  client.connect('localhost:50052', {
     plaintext: true
  });

  const data = {
    take: 'Bert',
    cards: [
      {type: "V", color: "Pique"},
      {type: "9", color: "Pique"},
      {type: "10", color: "Pique"},
      {type: "8", color: "Pique"},
      {type: "D", color: "Pique"},
      {type: "D", color: "Pique"},
    ],
    constraints: ["Passe"]
  };
  const serviceName = 'proto.gametake_history.v1.GameTakeHistory/Add'
  const response = client.invoke(serviceName, data);

  check(response, {
    'status is OK': (r) => r && r.status === grpc.StatusOK,
  });

  console.log(JSON.stringify(response.message));

  client.close();
  sleep(1);
};

