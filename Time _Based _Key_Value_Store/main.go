// https://leetcode.com/problems/time-based-key-value-store/

package main

type Data struct {
	val  string
	time int
}

type TimeMap struct {
	mapper map[string][]Data
}

func Constructor() TimeMap {
	return TimeMap{
		mapper: make(map[string][]Data),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	_, ok := this.mapper[key]
	if !ok {
		this.mapper[key] = make([]Data, 0)
	}
	this.mapper[key] = append(this.mapper[key], Data{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	_, ok := this.mapper[key]
	if !ok {
		return ""
	}
	return binarySearch(this.mapper[key], timestamp)
}

func binarySearch(dataList []Data, time int) string {
	low, high := 0, len(dataList)-1

	for low < high {
		mid := (low + high) / 2
		if dataList[mid].time == time {
			return dataList[mid].val
		}
		if dataList[mid].time < time {
			if dataList[mid+1].time > time {
				return dataList[mid].val
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if dataList[low].time <= time {
		return dataList[low].val
	} else {
		return ""
	}
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
