@extends('dashboard.layouts.main')

@section('container')
<div class="row">
  <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
      <h1 class="h2">ABSENSI TEKNISI MAINTENANCE</h1>
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
  
    <h2>Section title</h2>
    <div class="table-responsive">
      <table id="myTable" class="table table-striped w-100">
        <thead>
            <tr>
                <th>ID</th>
                <th>Name</th>
                <th>x</th>
                <th>y</th>
                <th>Status</th>
                <th>Status</th>
            </tr>          
        </thead>
        <tbody>
        </tbody>
    </table>
    </div>
      

</div>
@endsection

@push('script')
<script src="/js/form/absensiteknisi.js">
  

  
</script>
@endpush