import pandas as pd
from pathlib import Path

dataDir = "data"

p = Path(dataDir)


def catch(func, *args, handle=lambda e: e, **kwargs):
    try:
        return func(*args, **kwargs)
    except Exception as e:
        return handle(e)


data = [
    catch(pd.read_csv, item, encoding="cp932", skiprows=2, header=None)
    for item in p.iterdir()
    if item.is_file() and item.suffix == ".csv"
]

df = pd.concat([records for records in data if not isinstance(records, Exception)])

df = df[[3, 6]]
df = df.rename(columns={3: "date_of_use", 6: "amount"})
df["date_of_use"] = pd.to_datetime(df["date_of_use"], format="%Y年%m月%d日")
df["amount"] = df["amount"].str.replace(",", "").astype(int)
df = df.set_index("date_of_use")

print(df.resample("M").sum())
print("-" * 30)
print(df.resample("Y").sum())
