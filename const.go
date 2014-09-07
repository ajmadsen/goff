/* GENERATED BY genffconst.sh, DO NOT EDIT */

package ff

/*
#cgo pkg-config: libavformat libavcodec libavutil
#include <libavformat/avformat.h>
#include <libavcodec/avcodec.h>
#include <libavutil/avutil.h>
*/
import "C"

type CodecID int

var (
	AV_CODEC_ID_012V                        = CodecID(C.AV_CODEC_ID_012V)
	AV_CODEC_ID_4XM                         = CodecID(C.AV_CODEC_ID_4XM)
	AV_CODEC_ID_8BPS                        = CodecID(C.AV_CODEC_ID_8BPS)
	AV_CODEC_ID_8SVX_EXP                    = CodecID(C.AV_CODEC_ID_8SVX_EXP)
	AV_CODEC_ID_8SVX_FIB                    = CodecID(C.AV_CODEC_ID_8SVX_FIB)
	AV_CODEC_ID_A64_MULTI                   = CodecID(C.AV_CODEC_ID_A64_MULTI)
	AV_CODEC_ID_A64_MULTI5                  = CodecID(C.AV_CODEC_ID_A64_MULTI5)
	AV_CODEC_ID_AAC                         = CodecID(C.AV_CODEC_ID_AAC)
	AV_CODEC_ID_AAC_LATM                    = CodecID(C.AV_CODEC_ID_AAC_LATM)
	AV_CODEC_ID_AASC                        = CodecID(C.AV_CODEC_ID_AASC)
	AV_CODEC_ID_AC3                         = CodecID(C.AV_CODEC_ID_AC3)
	AV_CODEC_ID_ADPCM_4XM                   = CodecID(C.AV_CODEC_ID_ADPCM_4XM)
	AV_CODEC_ID_ADPCM_ADX                   = CodecID(C.AV_CODEC_ID_ADPCM_ADX)
	AV_CODEC_ID_ADPCM_AFC                   = CodecID(C.AV_CODEC_ID_ADPCM_AFC)
	AV_CODEC_ID_ADPCM_CT                    = CodecID(C.AV_CODEC_ID_ADPCM_CT)
	AV_CODEC_ID_ADPCM_DTK                   = CodecID(C.AV_CODEC_ID_ADPCM_DTK)
	AV_CODEC_ID_ADPCM_EA                    = CodecID(C.AV_CODEC_ID_ADPCM_EA)
	AV_CODEC_ID_ADPCM_EA_MAXIS_XA           = CodecID(C.AV_CODEC_ID_ADPCM_EA_MAXIS_XA)
	AV_CODEC_ID_ADPCM_EA_R1                 = CodecID(C.AV_CODEC_ID_ADPCM_EA_R1)
	AV_CODEC_ID_ADPCM_EA_R2                 = CodecID(C.AV_CODEC_ID_ADPCM_EA_R2)
	AV_CODEC_ID_ADPCM_EA_R3                 = CodecID(C.AV_CODEC_ID_ADPCM_EA_R3)
	AV_CODEC_ID_ADPCM_EA_XAS                = CodecID(C.AV_CODEC_ID_ADPCM_EA_XAS)
	AV_CODEC_ID_ADPCM_G722                  = CodecID(C.AV_CODEC_ID_ADPCM_G722)
	AV_CODEC_ID_ADPCM_G726                  = CodecID(C.AV_CODEC_ID_ADPCM_G726)
	AV_CODEC_ID_ADPCM_G726LE                = CodecID(C.AV_CODEC_ID_ADPCM_G726LE)
	AV_CODEC_ID_ADPCM_IMA_AMV               = CodecID(C.AV_CODEC_ID_ADPCM_IMA_AMV)
	AV_CODEC_ID_ADPCM_IMA_APC               = CodecID(C.AV_CODEC_ID_ADPCM_IMA_APC)
	AV_CODEC_ID_ADPCM_IMA_DK3               = CodecID(C.AV_CODEC_ID_ADPCM_IMA_DK3)
	AV_CODEC_ID_ADPCM_IMA_DK4               = CodecID(C.AV_CODEC_ID_ADPCM_IMA_DK4)
	AV_CODEC_ID_ADPCM_IMA_EA_EACS           = CodecID(C.AV_CODEC_ID_ADPCM_IMA_EA_EACS)
	AV_CODEC_ID_ADPCM_IMA_EA_SEAD           = CodecID(C.AV_CODEC_ID_ADPCM_IMA_EA_SEAD)
	AV_CODEC_ID_ADPCM_IMA_ISS               = CodecID(C.AV_CODEC_ID_ADPCM_IMA_ISS)
	AV_CODEC_ID_ADPCM_IMA_OKI               = CodecID(C.AV_CODEC_ID_ADPCM_IMA_OKI)
	AV_CODEC_ID_ADPCM_IMA_QT                = CodecID(C.AV_CODEC_ID_ADPCM_IMA_QT)
	AV_CODEC_ID_ADPCM_IMA_RAD               = CodecID(C.AV_CODEC_ID_ADPCM_IMA_RAD)
	AV_CODEC_ID_ADPCM_IMA_SMJPEG            = CodecID(C.AV_CODEC_ID_ADPCM_IMA_SMJPEG)
	AV_CODEC_ID_ADPCM_IMA_WAV               = CodecID(C.AV_CODEC_ID_ADPCM_IMA_WAV)
	AV_CODEC_ID_ADPCM_IMA_WS                = CodecID(C.AV_CODEC_ID_ADPCM_IMA_WS)
	AV_CODEC_ID_ADPCM_MS                    = CodecID(C.AV_CODEC_ID_ADPCM_MS)
	AV_CODEC_ID_ADPCM_SBPRO_2               = CodecID(C.AV_CODEC_ID_ADPCM_SBPRO_2)
	AV_CODEC_ID_ADPCM_SBPRO_3               = CodecID(C.AV_CODEC_ID_ADPCM_SBPRO_3)
	AV_CODEC_ID_ADPCM_SBPRO_4               = CodecID(C.AV_CODEC_ID_ADPCM_SBPRO_4)
	AV_CODEC_ID_ADPCM_SWF                   = CodecID(C.AV_CODEC_ID_ADPCM_SWF)
	AV_CODEC_ID_ADPCM_THP                   = CodecID(C.AV_CODEC_ID_ADPCM_THP)
	AV_CODEC_ID_ADPCM_VIMA                  = CodecID(C.AV_CODEC_ID_ADPCM_VIMA)
	AV_CODEC_ID_ADPCM_VIMA_DEPRECATED       = CodecID(C.AV_CODEC_ID_ADPCM_VIMA_DEPRECATED)
	AV_CODEC_ID_ADPCM_XA                    = CodecID(C.AV_CODEC_ID_ADPCM_XA)
	AV_CODEC_ID_ADPCM_YAMAHA                = CodecID(C.AV_CODEC_ID_ADPCM_YAMAHA)
	AV_CODEC_ID_AIC                         = CodecID(C.AV_CODEC_ID_AIC)
	AV_CODEC_ID_ALAC                        = CodecID(C.AV_CODEC_ID_ALAC)
	AV_CODEC_ID_ALIAS_PIX                   = CodecID(C.AV_CODEC_ID_ALIAS_PIX)
	AV_CODEC_ID_AMR_NB                      = CodecID(C.AV_CODEC_ID_AMR_NB)
	AV_CODEC_ID_AMR_WB                      = CodecID(C.AV_CODEC_ID_AMR_WB)
	AV_CODEC_ID_AMV                         = CodecID(C.AV_CODEC_ID_AMV)
	AV_CODEC_ID_ANM                         = CodecID(C.AV_CODEC_ID_ANM)
	AV_CODEC_ID_ANSI                        = CodecID(C.AV_CODEC_ID_ANSI)
	AV_CODEC_ID_APE                         = CodecID(C.AV_CODEC_ID_APE)
	AV_CODEC_ID_ASS                         = CodecID(C.AV_CODEC_ID_ASS)
	AV_CODEC_ID_ASV1                        = CodecID(C.AV_CODEC_ID_ASV1)
	AV_CODEC_ID_ASV2                        = CodecID(C.AV_CODEC_ID_ASV2)
	AV_CODEC_ID_ATRAC1                      = CodecID(C.AV_CODEC_ID_ATRAC1)
	AV_CODEC_ID_ATRAC3                      = CodecID(C.AV_CODEC_ID_ATRAC3)
	AV_CODEC_ID_ATRAC3P                     = CodecID(C.AV_CODEC_ID_ATRAC3P)
	AV_CODEC_ID_AURA                        = CodecID(C.AV_CODEC_ID_AURA)
	AV_CODEC_ID_AURA2                       = CodecID(C.AV_CODEC_ID_AURA2)
	AV_CODEC_ID_AVRN                        = CodecID(C.AV_CODEC_ID_AVRN)
	AV_CODEC_ID_AVRP                        = CodecID(C.AV_CODEC_ID_AVRP)
	AV_CODEC_ID_AVS                         = CodecID(C.AV_CODEC_ID_AVS)
	AV_CODEC_ID_AVUI                        = CodecID(C.AV_CODEC_ID_AVUI)
	AV_CODEC_ID_AYUV                        = CodecID(C.AV_CODEC_ID_AYUV)
	AV_CODEC_ID_BETHSOFTVID                 = CodecID(C.AV_CODEC_ID_BETHSOFTVID)
	AV_CODEC_ID_BFI                         = CodecID(C.AV_CODEC_ID_BFI)
	AV_CODEC_ID_BINKAUDIO_DCT               = CodecID(C.AV_CODEC_ID_BINKAUDIO_DCT)
	AV_CODEC_ID_BINKAUDIO_RDFT              = CodecID(C.AV_CODEC_ID_BINKAUDIO_RDFT)
	AV_CODEC_ID_BINKVIDEO                   = CodecID(C.AV_CODEC_ID_BINKVIDEO)
	AV_CODEC_ID_BINTEXT                     = CodecID(C.AV_CODEC_ID_BINTEXT)
	AV_CODEC_ID_BIN_DATA                    = CodecID(C.AV_CODEC_ID_BIN_DATA)
	AV_CODEC_ID_BMP                         = CodecID(C.AV_CODEC_ID_BMP)
	AV_CODEC_ID_BMV_AUDIO                   = CodecID(C.AV_CODEC_ID_BMV_AUDIO)
	AV_CODEC_ID_BMV_VIDEO                   = CodecID(C.AV_CODEC_ID_BMV_VIDEO)
	AV_CODEC_ID_BRENDER_PIX                 = CodecID(C.AV_CODEC_ID_BRENDER_PIX)
	AV_CODEC_ID_BRENDER_PIX_DEPRECATED      = CodecID(C.AV_CODEC_ID_BRENDER_PIX_DEPRECATED)
	AV_CODEC_ID_C93                         = CodecID(C.AV_CODEC_ID_C93)
	AV_CODEC_ID_CAVS                        = CodecID(C.AV_CODEC_ID_CAVS)
	AV_CODEC_ID_CDGRAPHICS                  = CodecID(C.AV_CODEC_ID_CDGRAPHICS)
	AV_CODEC_ID_CDXL                        = CodecID(C.AV_CODEC_ID_CDXL)
	AV_CODEC_ID_CELT                        = CodecID(C.AV_CODEC_ID_CELT)
	AV_CODEC_ID_CINEPAK                     = CodecID(C.AV_CODEC_ID_CINEPAK)
	AV_CODEC_ID_CLJR                        = CodecID(C.AV_CODEC_ID_CLJR)
	AV_CODEC_ID_CLLC                        = CodecID(C.AV_CODEC_ID_CLLC)
	AV_CODEC_ID_CMV                         = CodecID(C.AV_CODEC_ID_CMV)
	AV_CODEC_ID_COMFORT_NOISE               = CodecID(C.AV_CODEC_ID_COMFORT_NOISE)
	AV_CODEC_ID_COOK                        = CodecID(C.AV_CODEC_ID_COOK)
	AV_CODEC_ID_CPIA                        = CodecID(C.AV_CODEC_ID_CPIA)
	AV_CODEC_ID_CSCD                        = CodecID(C.AV_CODEC_ID_CSCD)
	AV_CODEC_ID_CYUV                        = CodecID(C.AV_CODEC_ID_CYUV)
	AV_CODEC_ID_DFA                         = CodecID(C.AV_CODEC_ID_DFA)
	AV_CODEC_ID_DIRAC                       = CodecID(C.AV_CODEC_ID_DIRAC)
	AV_CODEC_ID_DNXHD                       = CodecID(C.AV_CODEC_ID_DNXHD)
	AV_CODEC_ID_DPX                         = CodecID(C.AV_CODEC_ID_DPX)
	AV_CODEC_ID_DSD_LSBF                    = CodecID(C.AV_CODEC_ID_DSD_LSBF)
	AV_CODEC_ID_DSD_LSBF_PLANAR             = CodecID(C.AV_CODEC_ID_DSD_LSBF_PLANAR)
	AV_CODEC_ID_DSD_MSBF                    = CodecID(C.AV_CODEC_ID_DSD_MSBF)
	AV_CODEC_ID_DSD_MSBF_PLANAR             = CodecID(C.AV_CODEC_ID_DSD_MSBF_PLANAR)
	AV_CODEC_ID_DSICINAUDIO                 = CodecID(C.AV_CODEC_ID_DSICINAUDIO)
	AV_CODEC_ID_DSICINVIDEO                 = CodecID(C.AV_CODEC_ID_DSICINVIDEO)
	AV_CODEC_ID_DTS                         = CodecID(C.AV_CODEC_ID_DTS)
	AV_CODEC_ID_DVAUDIO                     = CodecID(C.AV_CODEC_ID_DVAUDIO)
	AV_CODEC_ID_DVB_SUBTITLE                = CodecID(C.AV_CODEC_ID_DVB_SUBTITLE)
	AV_CODEC_ID_DVB_TELETEXT                = CodecID(C.AV_CODEC_ID_DVB_TELETEXT)
	AV_CODEC_ID_DVD_NAV                     = CodecID(C.AV_CODEC_ID_DVD_NAV)
	AV_CODEC_ID_DVD_SUBTITLE                = CodecID(C.AV_CODEC_ID_DVD_SUBTITLE)
	AV_CODEC_ID_DVVIDEO                     = CodecID(C.AV_CODEC_ID_DVVIDEO)
	AV_CODEC_ID_DXA                         = CodecID(C.AV_CODEC_ID_DXA)
	AV_CODEC_ID_DXTORY                      = CodecID(C.AV_CODEC_ID_DXTORY)
	AV_CODEC_ID_EAC3                        = CodecID(C.AV_CODEC_ID_EAC3)
	AV_CODEC_ID_EIA_608                     = CodecID(C.AV_CODEC_ID_EIA_608)
	AV_CODEC_ID_ESCAPE124                   = CodecID(C.AV_CODEC_ID_ESCAPE124)
	AV_CODEC_ID_ESCAPE130                   = CodecID(C.AV_CODEC_ID_ESCAPE130)
	AV_CODEC_ID_ESCAPE130_DEPRECATED        = CodecID(C.AV_CODEC_ID_ESCAPE130_DEPRECATED)
	AV_CODEC_ID_EVRC                        = CodecID(C.AV_CODEC_ID_EVRC)
	AV_CODEC_ID_EXR                         = CodecID(C.AV_CODEC_ID_EXR)
	AV_CODEC_ID_EXR_DEPRECATED              = CodecID(C.AV_CODEC_ID_EXR_DEPRECATED)
	AV_CODEC_ID_FFMETADATA                  = CodecID(C.AV_CODEC_ID_FFMETADATA)
	AV_CODEC_ID_FFV1                        = CodecID(C.AV_CODEC_ID_FFV1)
	AV_CODEC_ID_FFVHUFF                     = CodecID(C.AV_CODEC_ID_FFVHUFF)
	AV_CODEC_ID_FFWAVESYNTH                 = CodecID(C.AV_CODEC_ID_FFWAVESYNTH)
	AV_CODEC_ID_FIC                         = CodecID(C.AV_CODEC_ID_FIC)
	AV_CODEC_ID_FIRST_AUDIO                 = CodecID(C.AV_CODEC_ID_FIRST_AUDIO)
	AV_CODEC_ID_FIRST_SUBTITLE              = CodecID(C.AV_CODEC_ID_FIRST_SUBTITLE)
	AV_CODEC_ID_FIRST_UNKNOWN               = CodecID(C.AV_CODEC_ID_FIRST_UNKNOWN)
	AV_CODEC_ID_FLAC                        = CodecID(C.AV_CODEC_ID_FLAC)
	AV_CODEC_ID_FLASHSV                     = CodecID(C.AV_CODEC_ID_FLASHSV)
	AV_CODEC_ID_FLASHSV2                    = CodecID(C.AV_CODEC_ID_FLASHSV2)
	AV_CODEC_ID_FLIC                        = CodecID(C.AV_CODEC_ID_FLIC)
	AV_CODEC_ID_FLV1                        = CodecID(C.AV_CODEC_ID_FLV1)
	AV_CODEC_ID_FRAPS                       = CodecID(C.AV_CODEC_ID_FRAPS)
	AV_CODEC_ID_FRWU                        = CodecID(C.AV_CODEC_ID_FRWU)
	AV_CODEC_ID_G2M                         = CodecID(C.AV_CODEC_ID_G2M)
	AV_CODEC_ID_G2M_DEPRECATED              = CodecID(C.AV_CODEC_ID_G2M_DEPRECATED)
	AV_CODEC_ID_G723_1                      = CodecID(C.AV_CODEC_ID_G723_1)
	AV_CODEC_ID_G729                        = CodecID(C.AV_CODEC_ID_G729)
	AV_CODEC_ID_GIF                         = CodecID(C.AV_CODEC_ID_GIF)
	AV_CODEC_ID_GSM                         = CodecID(C.AV_CODEC_ID_GSM)
	AV_CODEC_ID_GSM_MS                      = CodecID(C.AV_CODEC_ID_GSM_MS)
	AV_CODEC_ID_H261                        = CodecID(C.AV_CODEC_ID_H261)
	AV_CODEC_ID_H263                        = CodecID(C.AV_CODEC_ID_H263)
	AV_CODEC_ID_H263I                       = CodecID(C.AV_CODEC_ID_H263I)
	AV_CODEC_ID_H263P                       = CodecID(C.AV_CODEC_ID_H263P)
	AV_CODEC_ID_H264                        = CodecID(C.AV_CODEC_ID_H264)
	AV_CODEC_ID_HDMV_PGS_SUBTITLE           = CodecID(C.AV_CODEC_ID_HDMV_PGS_SUBTITLE)
	AV_CODEC_ID_HEVC                        = CodecID(C.AV_CODEC_ID_HEVC)
	AV_CODEC_ID_HEVC_DEPRECATED             = CodecID(C.AV_CODEC_ID_HEVC_DEPRECATED)
	AV_CODEC_ID_HNM4_VIDEO                  = CodecID(C.AV_CODEC_ID_HNM4_VIDEO)
	AV_CODEC_ID_HUFFYUV                     = CodecID(C.AV_CODEC_ID_HUFFYUV)
	AV_CODEC_ID_IAC                         = CodecID(C.AV_CODEC_ID_IAC)
	AV_CODEC_ID_IDCIN                       = CodecID(C.AV_CODEC_ID_IDCIN)
	AV_CODEC_ID_IDF                         = CodecID(C.AV_CODEC_ID_IDF)
	AV_CODEC_ID_IFF_BYTERUN1                = CodecID(C.AV_CODEC_ID_IFF_BYTERUN1)
	AV_CODEC_ID_IFF_ILBM                    = CodecID(C.AV_CODEC_ID_IFF_ILBM)
	AV_CODEC_ID_ILBC                        = CodecID(C.AV_CODEC_ID_ILBC)
	AV_CODEC_ID_IMC                         = CodecID(C.AV_CODEC_ID_IMC)
	AV_CODEC_ID_INDEO2                      = CodecID(C.AV_CODEC_ID_INDEO2)
	AV_CODEC_ID_INDEO3                      = CodecID(C.AV_CODEC_ID_INDEO3)
	AV_CODEC_ID_INDEO4                      = CodecID(C.AV_CODEC_ID_INDEO4)
	AV_CODEC_ID_INDEO5                      = CodecID(C.AV_CODEC_ID_INDEO5)
	AV_CODEC_ID_INTERPLAY_DPCM              = CodecID(C.AV_CODEC_ID_INTERPLAY_DPCM)
	AV_CODEC_ID_INTERPLAY_VIDEO             = CodecID(C.AV_CODEC_ID_INTERPLAY_VIDEO)
	AV_CODEC_ID_JACOSUB                     = CodecID(C.AV_CODEC_ID_JACOSUB)
	AV_CODEC_ID_JPEG2000                    = CodecID(C.AV_CODEC_ID_JPEG2000)
	AV_CODEC_ID_JPEGLS                      = CodecID(C.AV_CODEC_ID_JPEGLS)
	AV_CODEC_ID_JV                          = CodecID(C.AV_CODEC_ID_JV)
	AV_CODEC_ID_KGV1                        = CodecID(C.AV_CODEC_ID_KGV1)
	AV_CODEC_ID_KMVC                        = CodecID(C.AV_CODEC_ID_KMVC)
	AV_CODEC_ID_LAGARITH                    = CodecID(C.AV_CODEC_ID_LAGARITH)
	AV_CODEC_ID_LJPEG                       = CodecID(C.AV_CODEC_ID_LJPEG)
	AV_CODEC_ID_LOCO                        = CodecID(C.AV_CODEC_ID_LOCO)
	AV_CODEC_ID_MACE3                       = CodecID(C.AV_CODEC_ID_MACE3)
	AV_CODEC_ID_MACE6                       = CodecID(C.AV_CODEC_ID_MACE6)
	AV_CODEC_ID_MAD                         = CodecID(C.AV_CODEC_ID_MAD)
	AV_CODEC_ID_MDEC                        = CodecID(C.AV_CODEC_ID_MDEC)
	AV_CODEC_ID_METASOUND                   = CodecID(C.AV_CODEC_ID_METASOUND)
	AV_CODEC_ID_MICRODVD                    = CodecID(C.AV_CODEC_ID_MICRODVD)
	AV_CODEC_ID_MIMIC                       = CodecID(C.AV_CODEC_ID_MIMIC)
	AV_CODEC_ID_MJPEG                       = CodecID(C.AV_CODEC_ID_MJPEG)
	AV_CODEC_ID_MJPEGB                      = CodecID(C.AV_CODEC_ID_MJPEGB)
	AV_CODEC_ID_MLP                         = CodecID(C.AV_CODEC_ID_MLP)
	AV_CODEC_ID_MMVIDEO                     = CodecID(C.AV_CODEC_ID_MMVIDEO)
	AV_CODEC_ID_MOTIONPIXELS                = CodecID(C.AV_CODEC_ID_MOTIONPIXELS)
	AV_CODEC_ID_MOV_TEXT                    = CodecID(C.AV_CODEC_ID_MOV_TEXT)
	AV_CODEC_ID_MP1                         = CodecID(C.AV_CODEC_ID_MP1)
	AV_CODEC_ID_MP2                         = CodecID(C.AV_CODEC_ID_MP2)
	AV_CODEC_ID_MP3                         = CodecID(C.AV_CODEC_ID_MP3)
	AV_CODEC_ID_MP3ADU                      = CodecID(C.AV_CODEC_ID_MP3ADU)
	AV_CODEC_ID_MP3ON4                      = CodecID(C.AV_CODEC_ID_MP3ON4)
	AV_CODEC_ID_MP4ALS                      = CodecID(C.AV_CODEC_ID_MP4ALS)
	AV_CODEC_ID_MPEG1VIDEO                  = CodecID(C.AV_CODEC_ID_MPEG1VIDEO)
	AV_CODEC_ID_MPEG2TS                     = CodecID(C.AV_CODEC_ID_MPEG2TS)
	AV_CODEC_ID_MPEG2VIDEO                  = CodecID(C.AV_CODEC_ID_MPEG2VIDEO)
	AV_CODEC_ID_MPEG2VIDEO_XVMC             = CodecID(C.AV_CODEC_ID_MPEG2VIDEO_XVMC)
	AV_CODEC_ID_MPEG4                       = CodecID(C.AV_CODEC_ID_MPEG4)
	AV_CODEC_ID_MPEG4SYSTEMS                = CodecID(C.AV_CODEC_ID_MPEG4SYSTEMS)
	AV_CODEC_ID_MPL2                        = CodecID(C.AV_CODEC_ID_MPL2)
	AV_CODEC_ID_MSA1                        = CodecID(C.AV_CODEC_ID_MSA1)
	AV_CODEC_ID_MSMPEG4V1                   = CodecID(C.AV_CODEC_ID_MSMPEG4V1)
	AV_CODEC_ID_MSMPEG4V2                   = CodecID(C.AV_CODEC_ID_MSMPEG4V2)
	AV_CODEC_ID_MSMPEG4V3                   = CodecID(C.AV_CODEC_ID_MSMPEG4V3)
	AV_CODEC_ID_MSRLE                       = CodecID(C.AV_CODEC_ID_MSRLE)
	AV_CODEC_ID_MSS1                        = CodecID(C.AV_CODEC_ID_MSS1)
	AV_CODEC_ID_MSS2                        = CodecID(C.AV_CODEC_ID_MSS2)
	AV_CODEC_ID_MSVIDEO1                    = CodecID(C.AV_CODEC_ID_MSVIDEO1)
	AV_CODEC_ID_MSZH                        = CodecID(C.AV_CODEC_ID_MSZH)
	AV_CODEC_ID_MTS2                        = CodecID(C.AV_CODEC_ID_MTS2)
	AV_CODEC_ID_MUSEPACK7                   = CodecID(C.AV_CODEC_ID_MUSEPACK7)
	AV_CODEC_ID_MUSEPACK8                   = CodecID(C.AV_CODEC_ID_MUSEPACK8)
	AV_CODEC_ID_MVC1                        = CodecID(C.AV_CODEC_ID_MVC1)
	AV_CODEC_ID_MVC1_DEPRECATED             = CodecID(C.AV_CODEC_ID_MVC1_DEPRECATED)
	AV_CODEC_ID_MVC2                        = CodecID(C.AV_CODEC_ID_MVC2)
	AV_CODEC_ID_MVC2_DEPRECATED             = CodecID(C.AV_CODEC_ID_MVC2_DEPRECATED)
	AV_CODEC_ID_MXPEG                       = CodecID(C.AV_CODEC_ID_MXPEG)
	AV_CODEC_ID_NELLYMOSER                  = CodecID(C.AV_CODEC_ID_NELLYMOSER)
	AV_CODEC_ID_NONE                        = CodecID(C.AV_CODEC_ID_NONE)
	AV_CODEC_ID_NUV                         = CodecID(C.AV_CODEC_ID_NUV)
	AV_CODEC_ID_ON2AVC                      = CodecID(C.AV_CODEC_ID_ON2AVC)
	AV_CODEC_ID_OPUS                        = CodecID(C.AV_CODEC_ID_OPUS)
	AV_CODEC_ID_OPUS_DEPRECATED             = CodecID(C.AV_CODEC_ID_OPUS_DEPRECATED)
	AV_CODEC_ID_OTF                         = CodecID(C.AV_CODEC_ID_OTF)
	AV_CODEC_ID_PAF_AUDIO                   = CodecID(C.AV_CODEC_ID_PAF_AUDIO)
	AV_CODEC_ID_PAF_AUDIO_DEPRECATED        = CodecID(C.AV_CODEC_ID_PAF_AUDIO_DEPRECATED)
	AV_CODEC_ID_PAF_VIDEO                   = CodecID(C.AV_CODEC_ID_PAF_VIDEO)
	AV_CODEC_ID_PAF_VIDEO_DEPRECATED        = CodecID(C.AV_CODEC_ID_PAF_VIDEO_DEPRECATED)
	AV_CODEC_ID_PAM                         = CodecID(C.AV_CODEC_ID_PAM)
	AV_CODEC_ID_PBM                         = CodecID(C.AV_CODEC_ID_PBM)
	AV_CODEC_ID_PCM_ALAW                    = CodecID(C.AV_CODEC_ID_PCM_ALAW)
	AV_CODEC_ID_PCM_BLURAY                  = CodecID(C.AV_CODEC_ID_PCM_BLURAY)
	AV_CODEC_ID_PCM_DVD                     = CodecID(C.AV_CODEC_ID_PCM_DVD)
	AV_CODEC_ID_PCM_F32BE                   = CodecID(C.AV_CODEC_ID_PCM_F32BE)
	AV_CODEC_ID_PCM_F32LE                   = CodecID(C.AV_CODEC_ID_PCM_F32LE)
	AV_CODEC_ID_PCM_F64BE                   = CodecID(C.AV_CODEC_ID_PCM_F64BE)
	AV_CODEC_ID_PCM_F64LE                   = CodecID(C.AV_CODEC_ID_PCM_F64LE)
	AV_CODEC_ID_PCM_LXF                     = CodecID(C.AV_CODEC_ID_PCM_LXF)
	AV_CODEC_ID_PCM_MULAW                   = CodecID(C.AV_CODEC_ID_PCM_MULAW)
	AV_CODEC_ID_PCM_S16BE                   = CodecID(C.AV_CODEC_ID_PCM_S16BE)
	AV_CODEC_ID_PCM_S16BE_PLANAR            = CodecID(C.AV_CODEC_ID_PCM_S16BE_PLANAR)
	AV_CODEC_ID_PCM_S16LE                   = CodecID(C.AV_CODEC_ID_PCM_S16LE)
	AV_CODEC_ID_PCM_S16LE_PLANAR            = CodecID(C.AV_CODEC_ID_PCM_S16LE_PLANAR)
	AV_CODEC_ID_PCM_S24BE                   = CodecID(C.AV_CODEC_ID_PCM_S24BE)
	AV_CODEC_ID_PCM_S24DAUD                 = CodecID(C.AV_CODEC_ID_PCM_S24DAUD)
	AV_CODEC_ID_PCM_S24LE                   = CodecID(C.AV_CODEC_ID_PCM_S24LE)
	AV_CODEC_ID_PCM_S24LE_PLANAR            = CodecID(C.AV_CODEC_ID_PCM_S24LE_PLANAR)
	AV_CODEC_ID_PCM_S24LE_PLANAR_DEPRECATED = CodecID(C.AV_CODEC_ID_PCM_S24LE_PLANAR_DEPRECATED)
	AV_CODEC_ID_PCM_S32BE                   = CodecID(C.AV_CODEC_ID_PCM_S32BE)
	AV_CODEC_ID_PCM_S32LE                   = CodecID(C.AV_CODEC_ID_PCM_S32LE)
	AV_CODEC_ID_PCM_S32LE_PLANAR            = CodecID(C.AV_CODEC_ID_PCM_S32LE_PLANAR)
	AV_CODEC_ID_PCM_S32LE_PLANAR_DEPRECATED = CodecID(C.AV_CODEC_ID_PCM_S32LE_PLANAR_DEPRECATED)
	AV_CODEC_ID_PCM_S8                      = CodecID(C.AV_CODEC_ID_PCM_S8)
	AV_CODEC_ID_PCM_S8_PLANAR               = CodecID(C.AV_CODEC_ID_PCM_S8_PLANAR)
	AV_CODEC_ID_PCM_U16BE                   = CodecID(C.AV_CODEC_ID_PCM_U16BE)
	AV_CODEC_ID_PCM_U16LE                   = CodecID(C.AV_CODEC_ID_PCM_U16LE)
	AV_CODEC_ID_PCM_U24BE                   = CodecID(C.AV_CODEC_ID_PCM_U24BE)
	AV_CODEC_ID_PCM_U24LE                   = CodecID(C.AV_CODEC_ID_PCM_U24LE)
	AV_CODEC_ID_PCM_U32BE                   = CodecID(C.AV_CODEC_ID_PCM_U32BE)
	AV_CODEC_ID_PCM_U32LE                   = CodecID(C.AV_CODEC_ID_PCM_U32LE)
	AV_CODEC_ID_PCM_U8                      = CodecID(C.AV_CODEC_ID_PCM_U8)
	AV_CODEC_ID_PCM_ZORK                    = CodecID(C.AV_CODEC_ID_PCM_ZORK)
	AV_CODEC_ID_PCX                         = CodecID(C.AV_CODEC_ID_PCX)
	AV_CODEC_ID_PGM                         = CodecID(C.AV_CODEC_ID_PGM)
	AV_CODEC_ID_PGMYUV                      = CodecID(C.AV_CODEC_ID_PGMYUV)
	AV_CODEC_ID_PICTOR                      = CodecID(C.AV_CODEC_ID_PICTOR)
	AV_CODEC_ID_PJS                         = CodecID(C.AV_CODEC_ID_PJS)
	AV_CODEC_ID_PNG                         = CodecID(C.AV_CODEC_ID_PNG)
	AV_CODEC_ID_PPM                         = CodecID(C.AV_CODEC_ID_PPM)
	AV_CODEC_ID_PROBE                       = CodecID(C.AV_CODEC_ID_PROBE)
	AV_CODEC_ID_PRORES                      = CodecID(C.AV_CODEC_ID_PRORES)
	AV_CODEC_ID_PTX                         = CodecID(C.AV_CODEC_ID_PTX)
	AV_CODEC_ID_QCELP                       = CodecID(C.AV_CODEC_ID_QCELP)
	AV_CODEC_ID_QDM2                        = CodecID(C.AV_CODEC_ID_QDM2)
	AV_CODEC_ID_QDMC                        = CodecID(C.AV_CODEC_ID_QDMC)
	AV_CODEC_ID_QDRAW                       = CodecID(C.AV_CODEC_ID_QDRAW)
	AV_CODEC_ID_QPEG                        = CodecID(C.AV_CODEC_ID_QPEG)
	AV_CODEC_ID_QTRLE                       = CodecID(C.AV_CODEC_ID_QTRLE)
	AV_CODEC_ID_R10K                        = CodecID(C.AV_CODEC_ID_R10K)
	AV_CODEC_ID_R210                        = CodecID(C.AV_CODEC_ID_R210)
	AV_CODEC_ID_RALF                        = CodecID(C.AV_CODEC_ID_RALF)
	AV_CODEC_ID_RAWVIDEO                    = CodecID(C.AV_CODEC_ID_RAWVIDEO)
	AV_CODEC_ID_RA_144                      = CodecID(C.AV_CODEC_ID_RA_144)
	AV_CODEC_ID_RA_288                      = CodecID(C.AV_CODEC_ID_RA_288)
	AV_CODEC_ID_REALTEXT                    = CodecID(C.AV_CODEC_ID_REALTEXT)
	AV_CODEC_ID_RL2                         = CodecID(C.AV_CODEC_ID_RL2)
	AV_CODEC_ID_ROQ                         = CodecID(C.AV_CODEC_ID_ROQ)
	AV_CODEC_ID_ROQ_DPCM                    = CodecID(C.AV_CODEC_ID_ROQ_DPCM)
	AV_CODEC_ID_RPZA                        = CodecID(C.AV_CODEC_ID_RPZA)
	AV_CODEC_ID_RV10                        = CodecID(C.AV_CODEC_ID_RV10)
	AV_CODEC_ID_RV20                        = CodecID(C.AV_CODEC_ID_RV20)
	AV_CODEC_ID_RV30                        = CodecID(C.AV_CODEC_ID_RV30)
	AV_CODEC_ID_RV40                        = CodecID(C.AV_CODEC_ID_RV40)
	AV_CODEC_ID_S302M                       = CodecID(C.AV_CODEC_ID_S302M)
	AV_CODEC_ID_SAMI                        = CodecID(C.AV_CODEC_ID_SAMI)
	AV_CODEC_ID_SANM                        = CodecID(C.AV_CODEC_ID_SANM)
	AV_CODEC_ID_SANM_DEPRECATED             = CodecID(C.AV_CODEC_ID_SANM_DEPRECATED)
	AV_CODEC_ID_SGI                         = CodecID(C.AV_CODEC_ID_SGI)
	AV_CODEC_ID_SGIRLE                      = CodecID(C.AV_CODEC_ID_SGIRLE)
	AV_CODEC_ID_SGIRLE_DEPRECATED           = CodecID(C.AV_CODEC_ID_SGIRLE_DEPRECATED)
	AV_CODEC_ID_SHORTEN                     = CodecID(C.AV_CODEC_ID_SHORTEN)
	AV_CODEC_ID_SIPR                        = CodecID(C.AV_CODEC_ID_SIPR)
	AV_CODEC_ID_SMACKAUDIO                  = CodecID(C.AV_CODEC_ID_SMACKAUDIO)
	AV_CODEC_ID_SMACKVIDEO                  = CodecID(C.AV_CODEC_ID_SMACKVIDEO)
	AV_CODEC_ID_SMC                         = CodecID(C.AV_CODEC_ID_SMC)
	AV_CODEC_ID_SMPTE_KLV                   = CodecID(C.AV_CODEC_ID_SMPTE_KLV)
	AV_CODEC_ID_SMV                         = CodecID(C.AV_CODEC_ID_SMV)
	AV_CODEC_ID_SMVJPEG                     = CodecID(C.AV_CODEC_ID_SMVJPEG)
	AV_CODEC_ID_SNOW                        = CodecID(C.AV_CODEC_ID_SNOW)
	AV_CODEC_ID_SOL_DPCM                    = CodecID(C.AV_CODEC_ID_SOL_DPCM)
	AV_CODEC_ID_SONIC                       = CodecID(C.AV_CODEC_ID_SONIC)
	AV_CODEC_ID_SONIC_LS                    = CodecID(C.AV_CODEC_ID_SONIC_LS)
	AV_CODEC_ID_SP5X                        = CodecID(C.AV_CODEC_ID_SP5X)
	AV_CODEC_ID_SPEEX                       = CodecID(C.AV_CODEC_ID_SPEEX)
	AV_CODEC_ID_SRT                         = CodecID(C.AV_CODEC_ID_SRT)
	AV_CODEC_ID_SSA                         = CodecID(C.AV_CODEC_ID_SSA)
	AV_CODEC_ID_SUBRIP                      = CodecID(C.AV_CODEC_ID_SUBRIP)
	AV_CODEC_ID_SUBVIEWER                   = CodecID(C.AV_CODEC_ID_SUBVIEWER)
	AV_CODEC_ID_SUBVIEWER1                  = CodecID(C.AV_CODEC_ID_SUBVIEWER1)
	AV_CODEC_ID_SUNRAST                     = CodecID(C.AV_CODEC_ID_SUNRAST)
	AV_CODEC_ID_SVQ1                        = CodecID(C.AV_CODEC_ID_SVQ1)
	AV_CODEC_ID_SVQ3                        = CodecID(C.AV_CODEC_ID_SVQ3)
	AV_CODEC_ID_TAK                         = CodecID(C.AV_CODEC_ID_TAK)
	AV_CODEC_ID_TAK_DEPRECATED              = CodecID(C.AV_CODEC_ID_TAK_DEPRECATED)
	AV_CODEC_ID_TARGA                       = CodecID(C.AV_CODEC_ID_TARGA)
	AV_CODEC_ID_TARGA_Y216                  = CodecID(C.AV_CODEC_ID_TARGA_Y216)
	AV_CODEC_ID_TEXT                        = CodecID(C.AV_CODEC_ID_TEXT)
	AV_CODEC_ID_TGQ                         = CodecID(C.AV_CODEC_ID_TGQ)
	AV_CODEC_ID_TGV                         = CodecID(C.AV_CODEC_ID_TGV)
	AV_CODEC_ID_THEORA                      = CodecID(C.AV_CODEC_ID_THEORA)
	AV_CODEC_ID_THP                         = CodecID(C.AV_CODEC_ID_THP)
	AV_CODEC_ID_TIERTEXSEQVIDEO             = CodecID(C.AV_CODEC_ID_TIERTEXSEQVIDEO)
	AV_CODEC_ID_TIFF                        = CodecID(C.AV_CODEC_ID_TIFF)
	AV_CODEC_ID_TIMED_ID3                   = CodecID(C.AV_CODEC_ID_TIMED_ID3)
	AV_CODEC_ID_TMV                         = CodecID(C.AV_CODEC_ID_TMV)
	AV_CODEC_ID_TQI                         = CodecID(C.AV_CODEC_ID_TQI)
	AV_CODEC_ID_TRUEHD                      = CodecID(C.AV_CODEC_ID_TRUEHD)
	AV_CODEC_ID_TRUEMOTION1                 = CodecID(C.AV_CODEC_ID_TRUEMOTION1)
	AV_CODEC_ID_TRUEMOTION2                 = CodecID(C.AV_CODEC_ID_TRUEMOTION2)
	AV_CODEC_ID_TRUESPEECH                  = CodecID(C.AV_CODEC_ID_TRUESPEECH)
	AV_CODEC_ID_TSCC                        = CodecID(C.AV_CODEC_ID_TSCC)
	AV_CODEC_ID_TSCC2                       = CodecID(C.AV_CODEC_ID_TSCC2)
	AV_CODEC_ID_TTA                         = CodecID(C.AV_CODEC_ID_TTA)
	AV_CODEC_ID_TTF                         = CodecID(C.AV_CODEC_ID_TTF)
	AV_CODEC_ID_TWINVQ                      = CodecID(C.AV_CODEC_ID_TWINVQ)
	AV_CODEC_ID_TXD                         = CodecID(C.AV_CODEC_ID_TXD)
	AV_CODEC_ID_ULTI                        = CodecID(C.AV_CODEC_ID_ULTI)
	AV_CODEC_ID_UTVIDEO                     = CodecID(C.AV_CODEC_ID_UTVIDEO)
	AV_CODEC_ID_V210                        = CodecID(C.AV_CODEC_ID_V210)
	AV_CODEC_ID_V210X                       = CodecID(C.AV_CODEC_ID_V210X)
	AV_CODEC_ID_V308                        = CodecID(C.AV_CODEC_ID_V308)
	AV_CODEC_ID_V408                        = CodecID(C.AV_CODEC_ID_V408)
	AV_CODEC_ID_V410                        = CodecID(C.AV_CODEC_ID_V410)
	AV_CODEC_ID_VB                          = CodecID(C.AV_CODEC_ID_VB)
	AV_CODEC_ID_VBLE                        = CodecID(C.AV_CODEC_ID_VBLE)
	AV_CODEC_ID_VC1                         = CodecID(C.AV_CODEC_ID_VC1)
	AV_CODEC_ID_VC1IMAGE                    = CodecID(C.AV_CODEC_ID_VC1IMAGE)
	AV_CODEC_ID_VCR1                        = CodecID(C.AV_CODEC_ID_VCR1)
	AV_CODEC_ID_VIMA                        = CodecID(C.AV_CODEC_ID_VIMA)
	AV_CODEC_ID_VIXL                        = CodecID(C.AV_CODEC_ID_VIXL)
	AV_CODEC_ID_VMDAUDIO                    = CodecID(C.AV_CODEC_ID_VMDAUDIO)
	AV_CODEC_ID_VMDVIDEO                    = CodecID(C.AV_CODEC_ID_VMDVIDEO)
	AV_CODEC_ID_VMNC                        = CodecID(C.AV_CODEC_ID_VMNC)
	AV_CODEC_ID_VORBIS                      = CodecID(C.AV_CODEC_ID_VORBIS)
	AV_CODEC_ID_VOXWARE                     = CodecID(C.AV_CODEC_ID_VOXWARE)
	AV_CODEC_ID_VP3                         = CodecID(C.AV_CODEC_ID_VP3)
	AV_CODEC_ID_VP5                         = CodecID(C.AV_CODEC_ID_VP5)
	AV_CODEC_ID_VP6                         = CodecID(C.AV_CODEC_ID_VP6)
	AV_CODEC_ID_VP6A                        = CodecID(C.AV_CODEC_ID_VP6A)
	AV_CODEC_ID_VP6F                        = CodecID(C.AV_CODEC_ID_VP6F)
	AV_CODEC_ID_VP7                         = CodecID(C.AV_CODEC_ID_VP7)
	AV_CODEC_ID_VP7_DEPRECATED              = CodecID(C.AV_CODEC_ID_VP7_DEPRECATED)
	AV_CODEC_ID_VP8                         = CodecID(C.AV_CODEC_ID_VP8)
	AV_CODEC_ID_VP9                         = CodecID(C.AV_CODEC_ID_VP9)
	AV_CODEC_ID_VPLAYER                     = CodecID(C.AV_CODEC_ID_VPLAYER)
	AV_CODEC_ID_WAVPACK                     = CodecID(C.AV_CODEC_ID_WAVPACK)
	AV_CODEC_ID_WEBP                        = CodecID(C.AV_CODEC_ID_WEBP)
	AV_CODEC_ID_WEBP_DEPRECATED             = CodecID(C.AV_CODEC_ID_WEBP_DEPRECATED)
	AV_CODEC_ID_WEBVTT                      = CodecID(C.AV_CODEC_ID_WEBVTT)
	AV_CODEC_ID_WESTWOOD_SND1               = CodecID(C.AV_CODEC_ID_WESTWOOD_SND1)
	AV_CODEC_ID_WMALOSSLESS                 = CodecID(C.AV_CODEC_ID_WMALOSSLESS)
	AV_CODEC_ID_WMAPRO                      = CodecID(C.AV_CODEC_ID_WMAPRO)
	AV_CODEC_ID_WMAV1                       = CodecID(C.AV_CODEC_ID_WMAV1)
	AV_CODEC_ID_WMAV2                       = CodecID(C.AV_CODEC_ID_WMAV2)
	AV_CODEC_ID_WMAVOICE                    = CodecID(C.AV_CODEC_ID_WMAVOICE)
	AV_CODEC_ID_WMV1                        = CodecID(C.AV_CODEC_ID_WMV1)
	AV_CODEC_ID_WMV2                        = CodecID(C.AV_CODEC_ID_WMV2)
	AV_CODEC_ID_WMV3                        = CodecID(C.AV_CODEC_ID_WMV3)
	AV_CODEC_ID_WMV3IMAGE                   = CodecID(C.AV_CODEC_ID_WMV3IMAGE)
	AV_CODEC_ID_WNV1                        = CodecID(C.AV_CODEC_ID_WNV1)
	AV_CODEC_ID_WS_VQA                      = CodecID(C.AV_CODEC_ID_WS_VQA)
	AV_CODEC_ID_XAN_DPCM                    = CodecID(C.AV_CODEC_ID_XAN_DPCM)
	AV_CODEC_ID_XAN_WC3                     = CodecID(C.AV_CODEC_ID_XAN_WC3)
	AV_CODEC_ID_XAN_WC4                     = CodecID(C.AV_CODEC_ID_XAN_WC4)
	AV_CODEC_ID_XBIN                        = CodecID(C.AV_CODEC_ID_XBIN)
	AV_CODEC_ID_XBM                         = CodecID(C.AV_CODEC_ID_XBM)
	AV_CODEC_ID_XFACE                       = CodecID(C.AV_CODEC_ID_XFACE)
	AV_CODEC_ID_XSUB                        = CodecID(C.AV_CODEC_ID_XSUB)
	AV_CODEC_ID_XWD                         = CodecID(C.AV_CODEC_ID_XWD)
	AV_CODEC_ID_Y41P                        = CodecID(C.AV_CODEC_ID_Y41P)
	AV_CODEC_ID_YOP                         = CodecID(C.AV_CODEC_ID_YOP)
	AV_CODEC_ID_YUV4                        = CodecID(C.AV_CODEC_ID_YUV4)
	AV_CODEC_ID_ZEROCODEC                   = CodecID(C.AV_CODEC_ID_ZEROCODEC)
	AV_CODEC_ID_ZLIB                        = CodecID(C.AV_CODEC_ID_ZLIB)
	AV_CODEC_ID_ZMBV                        = CodecID(C.AV_CODEC_ID_ZMBV)
)
