function convertStringToArray(str) {
	var arr = [];
	var len, c;
	len = str.length;
	for(var i = 0;
		i < len;
		i++ ) {
		c = str.charCodeAt(i);
		if(c >= 0x010000 && c <= 0x10FFFF ) {
			arr.push(((c >> 18) & 0x07) | 0xF0);
			arr.push(((c >> 12) & 0x3F) | 0x80);
			arr.push(((c >> 6) & 0x3F) | 0x80);
			arr.push((c & 0x3F) | 0x80);
		} else if(c >= 0x000800 && c <= 0x00FFFF ) {
			arr.push(((c >> 12) & 0x0F) | 0xE0);
			arr.push(((c >> 6) & 0x3F) | 0x80);
			arr.push((c & 0x3F) | 0x80);
		} else if(c >= 0x000080 && c <= 0x0007FF ) {
			arr.push(((c >> 6) & 0x1F) | 0xC0);
			arr.push((c & 0x3F) | 0x80);
		} else {
			arr.push(c & 0xFF);
		}
	}
	return arr;
}

function convertArrayToString(arr){
	var byteArray = new Uint8Array(arr);
	var str = "", cc = 0, numBytes = 0;
	for(var i=0, len = byteArray.length; i<len; ++i){
		var v = byteArray[i];
		if(numBytes > 0){
			//2 bit determining that this is a tailing byte + 6 bit of payload
			if((cc&192) === 192){
				//processing tailing-bytes
				cc = (cc << 6) | (v & 63);
			}else{
				throw new Error("this is no tailing-byte");
			}
		}else if(v < 128){
			//single-byte
			numBytes = 1;
			cc = v;
		}else if(v < 192){
			//these are tailing-bytes
			throw new Error("invalid byte, this is a tailing-byte")
		}else if(v < 224){
			//3 bits of header + 5bits of payload
			numBytes = 2;
			cc = v & 31;
		}else if(v < 240){
			//4 bits of header + 4bit of payload
			numBytes = 3;
			cc = v & 15;
		}else{
			//UTF-8 theoretically supports up to 8 bytes containing up to 42bit of payload
			//but JS can only handle 16bit.
			throw new Error("invalid encoding, value out of range")
		}

		if(--numBytes === 0){
			str += String.fromCharCode(cc);
		}
	}
	if(numBytes){
		throw new Error("the bytes don't sum up");
	}
	return str;
}
