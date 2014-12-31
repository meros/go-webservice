define(['jquery', 'jquery.binarytransport', 'ProtoBuf', 'ByteBuffer'], 
       function($, bt, ProtoBuf, ByteBuffer) {
	   var builder = ProtoBuf.loadProto("message SumReq {\r\n\trequired uint32 a = 1;\r\n\trequired uint32 b = 2;\r\n}\r\n\r\nmessage SumResp {\r\n\trequired uint32 sum = 1;\r\n}\r\n");
	   
           var Root = builder.build();

	   // Req and Resp message classes
           var Req = Root.SumReq;
           var Resp = Root.SumResp;
	       
	   return {
	       createReq : function(initValue) {
		   // See https://github.com/dcodeIO/ProtoBuf.js/wiki/Builder
		   // on how to initialize messages
		   return new Req(initValue)
	       },	       
	       
	       // req: an instance of Req
	       // callback: a function taking an instance of Resp
	       call : function(req, callback){

		   // TODO: Lots of magic just to be able to send/receive binary body. 
		   // Encapsulate to enable easier bug fixing etc?
		   $.ajax({
                       url: "http://localhost:8080/sum/ws",
                       type: "POST",
                       data: req.toArrayBuffer(),
                       dataType: 'binary',
                       responseType: 'arraybuffer',
                       processData: false,
                       success: function(result) {
			   // TODO: check for error
                           var respBB = ByteBuffer.wrap(result)
                           var resp = new Resp.decode(respBB)
			   callback(resp)
                       },

                       error: function() {
			   // TODO: log error
                       }
                  });
	       }
	   };
       });
