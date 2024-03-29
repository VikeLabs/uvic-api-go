import time
from typing import List
from scripts.get_data.banner.schemas import BannerSection 
from scripts.get_data.banner.client import BannerClient
from scripts.get_data.sqlite.client import DB
import sys


def get_data():
    time_start = time.time()

    args = sys.argv
    banner = BannerClient()
    term: str
    if len(args) < 2:
        print(f"Get latest term", end=" ")
        t = banner.get_latest_term()
        print(f"\t[ok] found {t.code} - {t.description}")
        term = t.code
    else:
        term = args[1]


    print(f"\nSet term", end=" ")
    banner.term = term
    banner.set_term()
    print(f"\t\t[ok]")

    print("\nFetch data")
    offset = 0
    data: List[BannerSection] = list()
    while True:
        result = banner.get_data(offset)
        if result is None:
            break
        data = data + result
        offset += 1

    with DB(data) as db:
        # NOTE: don't change this insertion order
        # print("\nDrop existing data")
        # db.drop_existing_tables()
        # print("\t\t\t[ok]")

        print("\nSave to db")

        print("\t\t\t[pending] subjects", end=" ")
        db.save_subjects()
        print("\t[ok]")

        print("\t\t\t[pending] buildings", end=" ")
        db.save_buildings()
        print("\t[ok]")

        print("\t\t\t[pending] rooms", end=" ")
        db.save_rooms()
        print("\t[ok]")

        print("\t\t\t[pending] sessions", end=" ")
        db.save_sessions()
        print("\t[ok]")

        print("\t\t\t[ok]")

    time_end = time.time()
    print(f"\n[done]\t\t\ttook {(time_end - time_start) * 1000}ms")
