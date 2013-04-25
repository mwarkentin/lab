import envoy

r = envoy.run('./test.sh')
print r.status_code
print r.history
print r.std_out
