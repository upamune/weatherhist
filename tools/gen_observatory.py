import urllib.request
import lxml.html
import json

url = 'http://www.roy.hi-ho.ne.jp/ssai/mito_gis/report/gwoaddr.html'

html = urllib.request.urlopen(url).read().decode('shift-jis')
root = lxml.html.fromstring(html)

group_nums =         [ lxml.html.tostring(x, method='text', encoding='utf-8').decode('utf-8') for x in root.xpath('//tr/td[1]')]
obeservatory_ids =   [ lxml.html.tostring(x, method='text', encoding='utf-8').decode('utf-8') for x in  root.xpath('//tr/td[2]')]
obeservatory_names = [ lxml.html.tostring(x, method='text', encoding='utf-8').decode('utf-8') for x in root.xpath('//tr/td[3]')]

dict = {
    "obeservatories": []
}



for arr in zip(group_nums, obeservatory_ids, obeservatory_names):
    d = {
        "group_id": arr[0],
        "obeservatory_id": arr[1],
        "obeservatory_name": arr[2]
    }
    dict["obeservatories"].append(d)

print(dict)

f = open('obeservatories.json', 'w')
json.dump(dict, f, indent=2, ensure_ascii=False)
f.close()
