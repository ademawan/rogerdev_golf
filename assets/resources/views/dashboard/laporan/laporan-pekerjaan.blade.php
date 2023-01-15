@extends('dashboard.layouts.main')

@section('container')
<div class="row" id="container">
  <div class="d-flex justify-content-center flex-wrap flex-md-nowrap align-items-center border-bottom">
      <h1 class="h2 d-blockr">DAFTAR PERMINTAAN PERBAIKAN</h1>
      {{-- <div class="btn-toolbar mb-2 mb-md-0">
        <div class="btn-group me-2">
          <button type="button" class="btn btn-sm btn-outline-secondary">Share</button>
          <button type="button" class="btn btn-sm btn-outline-secondary">Export</button>
        </div>
        <button type="button" class="btn btn-sm btn-outline-secondary dropdown-toggle">
          <span data-feather="calendar"></span>
          This week
        </button>
      </div> --}}
    </div>
    <div class="row my-3">
      <div class="col-md-12">
        <div hidden>
          <button type="button" id="tombolRefreshTable" class="btn linka d-inline mb-2" style="background-color: rgba(104, 91, 19, 0.782) " >
            Refresh Table
          </button>
  
        </div>
      </div>
        
      </div>
      
    </div>
    <div class="row input-daterange mb-3">
      <div class="col-md-2">
          <input type="text" name="from_date" id="from_date" class="form-control" placeholder="From Date"
              readonly />
      </div>
      <div class="col-md-2">
          <input type="text" name="to_date" id="to_date" class="form-control" placeholder="To Date"
              readonly />
      </div>
      <div class="col-md-2">
          <button type="button" name="filter" id="filter" class="btn btn-primary">Filter</button>
          <button type="button" name="refresh" id="refresh" class="btn btn-default">Refresh</button>
      </div>
  </div>

    <div class="table-responsive">
      <table id="myTable" class="table table-striped" width="100%">
        <thead>
          <tr>
            <th>Nama Mesin</th>
            <th>Nama Teknisi</th>
            <th>Nama User</th>  
            <th>Tipe Perbaikan</th>  
            <th>Masalah Kerusakan</th>  
            <th>Kronologis Kejadian</th>  
            <th>Analisa</th>  
            <th>Tindakan</th>  
            <th>Progress</th>            
            <th>Keterangan</th>
            <th>Mulai Terjadi</th>
            <th>Mulai Perbaikan</th>
            <th>Selesai Perbaikan</th>
            <th>Pemberi Order</th>
          </tr>     
        </thead>
        <tbody>
          
        </tbody>
      </table>
    </div>

</div>
    
@endsection

@push('script')
<script src="/js/laporan/laporanpekerjaan.js"></script>
    
@endpush