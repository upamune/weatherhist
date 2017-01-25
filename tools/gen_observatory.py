import urllib.request
import lxml.html
import json

url = 'http://www.roy.hi-ho.ne.jp/ssai/mito_gis/report/gwoaddr.html'

html = urllib.request.urlopen(url).read().decode('shift-jis')
root = lxml.html.fromstring(html)

group_nums =         [ lxml.html.tostring(x, method='text', encoding='utf-8').decode('utf-8') for x in root.xpath('//tr/td[1]')]
observatory_ids =   [ lxml.html.tostring(x, method='text', encoding='utf-8').decode('utf-8') for x in  root.xpath('//tr/td[2]')]
observatory_names = [ lxml.html.tostring(x, method='text', encoding='utf-8').decode('utf-8') for x in root.xpath('//tr/td[3]')]

dict = {
    "observatories": []
}



for arr in zip(group_nums, observatory_ids, observatory_names):
    d = {
        "group_id": arr[0],
        "observatory_id": int(arr[1]),
        "observatory_name": arr[2]
    }
    dict["observatories"].append(d)

print(dict)

f = open('observatories.json', 'w')
json.dump(dict, f, indent=2, ensure_ascii=False)
f.close()
