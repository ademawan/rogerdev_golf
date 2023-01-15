$(document).ready( function () {
    
    $('#myTable').DataTable( {
      
      processing: true,
      paging:true,
      scrollX: true,
      scrollY: true,
      
      ajax: '',
      columns: [
              { data: 'name', name: 'name', class:'text-center' },
              { data: 'z', name: 'z' },
              { data: 'category', name: 'category', orderable:false },
              { data: 'b', name: 'b', orderable:false },
              { data: 'c', name: 'c', orderable:false },
              { data: 'd', name: 'd', orderable:false },
            
          ],
    });

   

  } );