import xlrd
import sys
import os
import json

station_file_name = "obs_stations.xlsx"

if os.path.exists(station_file_name):
    sheet = xlrd.open_workbook(station_file_name).sheet_by_index(0)
    dict = {
        "observatories": []
    }
    for r in range(sheet.nrows):
        if r in [0,1]:
            continue
        observatory_id = int(sheet.cell(r, 0).value)
        observatory_name = sheet.cell(r, 1).value
        group_id = int(sheet.cell(r, 3).value)
        observatory_type = sheet.cell(r, 4).value
        d = {
            "group_id": group_id,
            "observatory_id": observatory_id,
            "observatory_name": observatory_name,
            "observatory_type": observatory_type
        }
        dict["observatories"].append(d)

    f = open('observatories.json', 'w')
    json.dump(dict, f, indent=2, ensure_ascii=False)
    f.close()
else:
    sys.exit(1)

sys.exit(0)
