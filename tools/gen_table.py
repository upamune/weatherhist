import json

f = open('stations.json')
dict = json.load(f)
f.close()

tables = ["|station_id|group_id|name|type|", "|---:|---:|:--:|:--:|"]
row = "|%s|%s|%s|%s|"

stations = dict['stations']
for s in stations:
    r = row % (s['station_id'], s['group_id'], s['station_name'], s['station_type'])
    tables.append(r)

md = '\n'.join(tables)

md_file_name = "station.md"

with open(md_file_name, 'w') as f:
    f.write(md)

