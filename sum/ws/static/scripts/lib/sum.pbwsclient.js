define(['jquery', 'jquery.binarytransport', 'ProtoBuf', 'ByteBuffer'], 
       function($, bt, ProtoBuf, ByteBuffer) {
	   var builder = ProtoBuf.loadProto("message SumReq {\r\n\trequired uint32 a = 1;\r\n\trequired uint32 b = 2;\r\n}\r\n\r\nmessage SumResp {\r\n\trequired uint32 sum = 1;\r\n}\r\n");
	   
           var Root = builder.build();
           var Req = Root.SumReq;
           var Resp = Root.SumResp;
	       
	   return {
	       createReq : function(initValue) {
		   return new Req(initValue)
	       },	       
	       call : function(req, callback){
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
