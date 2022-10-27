import http from 'k6/http';

export const options = {
  discardResponseBodies: true,
  scenarios: {
    contacts: {
      executor: 'per-vu-iterations',
      vus: 100,
      maxDuration: '150s',
    },
  },
};

export default function () {
  const url = 'http://ec2-184-72-88-165.compute-1.amazonaws.com:8080/listarFilmes';
  const payload = JSON.stringify({
    email: 'aaa',
    password: 'bbb',
  });

  const params = {
    timeout: '150s',
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res  = http.get(url, params);
  // console.log(JSON.stringify(res.body));

}