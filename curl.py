import requests
import threading
import argparse
import time

# 创建一个锁对象
lock = threading.Lock()

success_time = 0


def curl_url(url):
    # 请求对应url,查看是否成功
    global success_time
    response = requests.get(url)
    if response.status_code == 200:
        # 上锁防止多线程同时修改
        with lock:
            success_time += 1


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("--qps", type=int, help="qps")
    parser.add_argument("--duration", type=int, help="duration")
    args = parser.parse_args()
    qps = args.qps
    duration = args.duration
    total_time = qps * duration
    url = "http://127.0.0.1:8888/api/student?name=许&page=1&page_size=20"
    threads = []
    for i in range(duration):
        time_start = time.time()
        for i in range(qps):
            t = threading.Thread(target=curl_url, args=(url,))
            threads.append(t)
            t.start()
        for t in threads:
            t.join()
        # 睡到距离start1s之后
        time_end = time.time()
        if time_end - time_start < 1:
            time.sleep(1 - (time_end - time_start))
        else:
            print("请求时间过长")
            exit(1)

    print("success_time:", success_time)
    print(f"成功率为{success_time/total_time}")
