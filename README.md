# Go-for-pastime  
This repository is NOT useful.  
There are datas and source codes for class report. 
I usually commit to master, because it is troublesome to merging and branching. 

## RTSP
RTSP is for my report.  
I do not have any responsibility when you submit your report using these codes.

### fConversion
This directory includes 2 Go codes to convert frequency.   
- down.go  
down.go is decimation.
This is not the punishment for ancient Roman soldiers.
This code downsamples the wave from 48kHz to 16 kHz.   
- up.go  
up.go is interpolation.
This code upsamples the wave from 8kHz to 32 kHz.  

### fftWav
Go codes of this directory do frequency analysis of waves by FFT.  

### psk
- bpsk  
This code detects BPSK signal.  
- qpsk   
This code detect QPSK signal.
I use the stereo audio as QPSK signal.
There are 2 BPSK signals for each left and right channels of streo audio.
Therefore, this code uses my BPSK detection twice.

## Import packages
### RTSP  
- [github.com/oov/audio/wave](github.com/oov/audio/wave)  
- [gonum.org/v1/plot](gonum.org/v1/plot)  
- [github.com/mjibson/go-dsp/fft](github.com/mjibson/go-dsp/fft)  
  
## Reference
1.  [https://godoc.org/github.com/oov/audio/wave](https://godoc.org/github.com/oov/audio/wave)  
2.  [https://godoc.org/gonum.org/v1/plot](https://godoc.org/gonum.org/v1/plot)  
3.  [https://godoc.org/github.com/mjibson/go-dsp/fft](https://godoc.org/github.com/mjibson/go-dsp/fft)  


