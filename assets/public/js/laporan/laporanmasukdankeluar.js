$(document).ready( function () {
    
    $('#myTable').DataTable( {
      dom: "Bfrtip",
buttons: [
  {
      extend: "excel",                // Extend the excel button
      insertCells: [                  // Add an insertCells config option 
          {
              cells: 'sCh',               // Target the header with smart selection
              content: 'New column C',    // New content for the cells
              pushCol: true,              // pushCol causes the column to be inserted
          },
          {
              cells: 'sC1:C-0',           // Target the data
              content: '',                // Add empty content
              pushCol: true               // push the columns to the right over one
          },
          {
              cells: 's5',              // Target data row 5 and 6
              content: '',                // Add empty content
              pushRow: true               // push the rows down to insert the content
          },
        
      ],
      excelStyles: {
          template: 'cyan_medium',    // Add a template to the result
      },
      exportOptions: { columns: [ ':visible' ] }
  },
],

      processing: true,
      paging:true,
      scrollX: true,
      scrollY: true,
      buttons: [
             'copy', 'print', 'pdf', 'excel'
               ],
      ajax: '',
      columns: [
              { data: 'name', name: 'name', class:'text-center' },
              { data: 'z', name: 'z' },
              { data: 'category', name: 'category', orderable:false },
              { data: 'b', name: 'b', orderable:false },
              { data: 'c', name: 'c', orderable:false },
            
          ],
    });

   

  } );