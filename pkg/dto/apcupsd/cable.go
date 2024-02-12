package apcupsd

type Cable uint8

const (
	CABLE__NO_CABLE      = Cable(0)
	CABLE__CUSTOM_SIMPLE = Cable(1)
	CABLE__APC_940_0119A = Cable(2)
	CABLE__APC_940_0127A = Cable(3)
	CABLE__APC_940_0128A = Cable(4)
	CABLE__APC_940_0020B = Cable(5)
	CABLE__APC_940_0020C = Cable(6)
	CABLE__APC_940_0023A = Cable(7)
	CABLE__MAM           = Cable(8)
	CABLE__APC_940_0095A = Cable(9)
	CABLE__APC_940_0095B = Cable(10)
	CABLE__APC_940_0095C = Cable(11)
	CABLE__CUSTOM_SMART  = Cable(12)
	CABLE__APC_940_0024B = Cable(121)
	CABLE__APC_940_0024C = Cable(122)
	CABLE__APC_940_1524C = Cable(123)
	CABLE__APC_940_0024G = Cable(124)
	CABLE__APC_940_0625A = Cable(125)
	CABLE__ETHERNET      = Cable(13)
	CABLE__USB           = Cable(14)
)
