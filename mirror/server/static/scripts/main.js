requirejs.config({
    baseUrl: 'scripts/lib',

    paths: {
        'ByteBuffer': 'ByteBufferAB.min',
	'ProtoBuf': 'ProtoBuf.min',
	'Long': 'Long.min',
	'jquery': 'jquery-2.1.3.min'
    },

    shim: {
	'jquery.binarytransport': {
            deps: ['jquery'],
            exports: 'BinaryTransport'
	}
    }
});

requirejs(['jquery', 'mirror.pbwsclient'],
	  function   ($, sum) {
              $("#calculateMirror").click(function(ev) {
		  // Create a req message
		  // Note: any value errors are caught here
		  // TODO: catch any errors and do something useful
                  var sumReq = sum.createReq({
                      "a": $("#valueA").val(),
                  })

		  // Call web service with callback function
		  // TODO: handle errors
		  // TODO: promises etc? Whats the new stuff in js land?
                  sum.call(sumReq, function(sumResp) {
		      $("#valueResult").html(sumResp.a)
		  });                
              });
          });
