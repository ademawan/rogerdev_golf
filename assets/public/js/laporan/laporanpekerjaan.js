$(function () {

    
    $.ajaxSetup({
        headers: {
            'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
        }
      });

      load_data();
    
    function load_data(from_date = '', to_date = '') {
        $('#myTable').DataTable({
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
      serverSide: true,
      paging:true,
      scrollX: true,
      scrollY: true,
      buttons: [
             'copy', 'print', 'pdf', 'excel'
               ],
      ajax: {
          url:'/dashboard/laporan/laporan-pekerjaan',
          type:'GET',
          data:{from_date:from_date, to_date:to_date}
      },
      columns: [
        { data: 'jenis_mesin', name: 'jenis_mesin', class:'text-center' },
        { data: 'teknisi', name: 'teknisi' },
        { data: 'user', name: 'user', orderable:false },
        { data: 'tipe_perbaikan', name: 'tipe_perbaikan', orderable:false },
        { data: 'masalah_kerusakan', name: 'masalah_kerusakan', orderable:false },
        { data: 'kronologis_kejadian', name: 'kronologis_kejadian', orderable:false },
        { data: 'analisa', name: 'analisa', orderable:false },
        { data: 'tindakan', name: 'tindakan', orderable:false },
        { data: 'progress', name: 'progress', orderable:false },
        { data: 'keterangan', name: 'keterangan', orderable:false },
        { data: 'mulai_terjadi', name: 'mulai_terjadi', orderable:false },
        { data: 'mulai_perbaikan', name: 'mulai_perbaikan', orderable:false },
        { data: 'selesai_perbaikan', name: 'selesai_perbaikan', orderable:false },
        { data: 'pemberi_order', name: 'pemberi_order', orderable:false },
            
          ]
        });
        }
   

    $('#tombolRefreshTable').click(function(){
        $('#myTable').DataTable().draw();
      });
      
      //datepicker
      
      $('.input-daterange').datepicker({

        todayBtn:'linked',
        format:'yyyy-mm-dd',
        autoclose:true,
        });

       
        //filter
        $('#filter').click(function () {
            var from_date = $('#from_date').val(); 
            var to_date = $('#to_date').val(); 

     
            if (from_date != '' && to_date != '') {
                $('#myTable').DataTable().destroy();
                load_data(from_date, to_date);
            } else {
                alert('Both Date is required');
            }
        });

        //refresh
        $('#refresh').click(function () {
            $('#from_date').val('');
            $('#to_date').val('');
            $('#myTable').DataTable().destroy();
            load_data();
        });

  } );