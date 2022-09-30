package ssTable

/*

索引是从数据区开始！
0 ─────────────────────────────────────────────────────────►
◄───────────────────────────
          dataLen          ◄──────────────────
                                indexLen     ◄──────────────┐
┌──────────────────────────┬─────────────────┬──────────────┤
│                          │                 │              │
│          数据区           │   稀疏索引区      │    元数据     │
│                          │                 │              │
└──────────────────────────┴─────────────────┴──────────────┘
*/

// MetaInfo 是 SSTable 的元数据，
// 元数据出现在磁盘文件的末尾
type MetaInfo struct {
	// 版本号
	version int64
	// 数据区起始索引
	dataStart int64
	// 数据区长度
	dataLen int64
	// 稀疏索引区起始索引
	indexStart int64
	// 稀疏索引区长度
	indexLen int64
}

func NewMetaInfo(dataLen, indexStart, indexLen int64) MetaInfo {
	return MetaInfo{
		version:    0,
		dataStart:  0,
		dataLen:    dataLen,
		indexStart: indexStart,
		indexLen:   indexLen,
	}
}
