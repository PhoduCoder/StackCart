Rule for comptability between method types and interface types and function receivers

Constructor returns |	Get receiver |	Interface assignment |	Method call |
MemoryCache	        |   Value	     |  Works	             |   Works      |
*MemoryCache	        Value	        Works	                Works
MemoryCache	            Pointer	        Requires &memory	    Works
*MemoryCache	        Pointer	        Works	                Works