var Client = {
	createNew: function (handler) {
		var res = {};
		res._handler = handler;
		res._inmsg = false;
		res._totalSize = 0;
		res._msgid = 0;
		res._dataArray = new ArrayBuffer(0);
		res._onmessage = function (evt) {
			var oldBuffer = res._dataArray;
			res._dataArray = new ArrayBuffer(
				oldBuffer.byteLength + evt.data.byteLength
			);
			var totalBufferView = new Uint8Array(res._dataArray);
			var oldBufferView = new Uint8Array(oldBuffer);
			var newBufferView = new Uint8Array(evt.data);
			totalBufferView.set(oldBufferView);
			totalBufferView.set(newBufferView, oldBuffer.byteLength);
			while (true) {
				if (res._inmsg === false) {
					if (res._dataArray.byteLength < 6) {
						break;
					}
					res._inmsg = true;
					var dataview = new DataView(res._dataArray.slice(0, 6));
					res._totalSize = dataview.getUint32(0);
					res._msgid = dataview.getUint16(4);
				} else {
					if (res._dataArray.byteLength >= res._totalSize) {
						var data = convertArrayToString(
							res._dataArray.slice(6, res._totalSize)
						);
						var topobj = JSON.parse(data);
						res._dataArray = res._dataArray.slice(res._totalSize);
						res._inmsg = false;
						if (res._handler.onmessage != null) {
							res._handler.onmessage(
								topobj.MsgName,
								JSON.parse(Base64.decode(topobj.Data))
							);
						}
					} else {
						break;
					}
				}
			}
		};
		res.send = function (topic, obj) {
			var toplayer = {
				MsgName: topic,
				Data: Base64.encode(JSON.stringify(obj)),
			};
			console.log(obj);
			var byteArray = new Uint8Array(
				convertStringToArray(JSON.stringify(toplayer))
			);
			var buffer = new ArrayBuffer(4 + 2 + byteArray.byteLength);
			var dataview = new DataView(buffer);
			dataview.setUint32(0, 4 + 2 + byteArray.byteLength);
			dataview.setUint16(4, 0);
			var head2Array = new Uint8Array(buffer, 4 + 2);
			head2Array.set(byteArray);
			res.wsconn.send(buffer);
		};
		res.dial = function (url) {
			res.wsconn = new WebSocket(url);
			res.wsconn.binaryType = "arraybuffer";
			res.wsconn.onopen = function (evt) {
				if (res._handler.onopen != null) {
					return res._handler.onopen(evt);
				}
			};
			res.wsconn.onclose = function (evt) {
				if (res._handler.onclose != null) {
					return res._handler.onclose(evt);
				}
			};
			res.wsconn.onmessage = res._onmessage;
			res.wsconn.onerror = function (evt) {
				if (res._handler.onerror != null) {
					return res._handler.onerror(evt);
				}
			};
		};
		return res;
	},
};
