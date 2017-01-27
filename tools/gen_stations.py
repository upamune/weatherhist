import xlrd
import sys
import os
import json

station_file_name = "obs_stations.xlsx"

if os.path.exists(station_file_name):
    sheet = xlrd.open_workbook(station_file_name).sheet_by_index(0)
    dict = {
        "stations": []
    }
    for r in range(sheet.nrows):
        if r in [0,1]:
            continue
        station_id = int(sheet.cell(r, 0).value)
        station_name = sheet.cell(r, 1).value
        group_id = int(sheet.cell(r, 3).value)
        station_type = sheet.cell(r, 4).value
        d = {
            "group_id": group_id,
            "station_id": station_id,
            "station_name": station_name,
            "station_type": station_type
        }
        dict["stations"].append(d)

    f = open('stations.json', 'w')
    json.dump(dict, f, indent=2, ensure_ascii=False)
    f.close()
else:
    sys.exit(1)

sys.exit(0)
