import sensor, image, time, ustruct
from pyb import USB_VCP

usb = USB_VCP()
sensor.reset()
sensor.set_pixformat(sensor.RGB565)
sensor.set_framesize(sensor.VGA)
sensor.skip_frames(time=2000)
header = [0x55, 0xAA, 0x5A, 0xA5]

while (True):
    img = sensor.snapshot().compress()
    for i in header:
        usb.send(i)
    usb.send(ustruct.pack("<L", img.size()))
    usb.send(img)
