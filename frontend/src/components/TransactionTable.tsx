type Txn = {
  timestamp: number;
  name: string;
  type: string;
  amount: number;
  status: string;
  description: string;
};

type TransactionTableProps = {
  transactions: Txn[];
  loading: boolean;
  page: number;
  total: number;
  perPage: number;
  sortBy: string;
  order: 'asc' | 'desc';
  onPageChange: (page: number) => void;
  onSortChange: (sortBy: string, order: 'asc' | 'desc') => void;
};

export default function TransactionTable({
  transactions,
  loading,
  page,
  total,
  perPage,
  sortBy,
  order,
  onPageChange,
  onSortChange,
}: TransactionTableProps) {
  const totalPages = Math.ceil(total / perPage) || 1;
  const currentPage = Math.min(page, totalPages);
  
  const handleSort = (column: string) => {
    if (sortBy === column) {
      // Toggle order if clicking the same column
      const newOrder = order === "asc" ? "desc" : "asc";
      onSortChange(column, newOrder);
    } else {
      // Default to ascending order when changing sort column
      onSortChange(column, "asc");
    }
    // Reset to first page when changing sort
    onPageChange(1);
  };

  const formatDate = (timestamp: number) => {
    return new Date(timestamp * 1000).toLocaleString("id-ID", {
      day: "2-digit",
      month: "short",
      year: "numeric",
      hour: "2-digit",
      minute: "2-digit",
    });
  };

  const formatCurrency = (amount: number) => {
    return new Intl.NumberFormat("id-ID", {
      style: "currency",
      currency: "IDR",
      minimumFractionDigits: 0,
    }).format(amount);
  };


  return (
    <div>
      <div className="card-header">
        <h2 className="card-title">Daftar Transaksi</h2>
      </div>

      {loading ? (
        <div className="text-center py-8">
          <p>Memuat data transaksi...</p>
        </div>
      ) : transactions.length === 0 ? (
        <div className="text-center py-8">
          <p>Tidak ada data transaksi</p>
        </div>
      ) : (
        <>
          <div className="table-container">
            <table className="table">
              <thead>
                <tr>
                  <th
                    onClick={() => handleSort("timestamp")}
                    className="cursor-pointer hover:text-primary"
                  >
                    <div className="flex items-center gap-1">
                      Tanggal
                      <span>{sortBy === "timestamp" && (order === "asc" ? "↑" : "↓")}</span>
                    </div>
                  </th>
                  <th>Nama</th>
                  <th>Tipe</th>
                  <th 
                    onClick={() => handleSort("amount")}
                    className="cursor-pointer hover:text-primary">
                      <div className="flex items-center gap-1">
                        Jumlah
                        <span>{sortBy === "amount" && (order === "asc" ? "↑" : "↓")}</span>
                      </div>
                  </th>
                  <th>Status</th>
                </tr>
              </thead>
              <tbody>
                {transactions.map((txn, index) => (
                  <tr key={`${txn.timestamp}-${index}`} className="hover:bg-gray-50">
                    <td className="text-sm">{formatDate(txn.timestamp)}</td>
                    <td>
                      <div className="font-medium">{txn.name}</div>
                      <div className="text-xs text-gray-500">{txn.description}</div>
                    </td>
                    <td className="capitalize">{txn.type.toLowerCase()}</td>
                    <td className="text-right font-medium">
                      {formatCurrency(txn.amount)}
                    </td>
                    <td>
                      <StatusBadge status={txn.status} />
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>

          
            <div className="pagination">
              <button
                onClick={() => onPageChange(Math.max(1, page - 1))}
                disabled={page === 1}
                className="pagination-btn"
              >
                Sebelumnya
              </button>
              <span className="text-sm">
                Halaman {currentPage} dari {totalPages}
              </span>
              <button
                onClick={() => onPageChange(page + 1)}
                disabled={currentPage >= totalPages}
                className="pagination-btn"
              >
                Selanjutnya
              </button>
            </div>
          
        </>
      )}
    </div>
  );
}

function StatusBadge({ status }: { status: string }) {
  if (status === "PENDING") return <span className="badge warning">⚠️ Pending</span>;
  if (status === "FAILED") return <span className="badge danger">❌ Failed</span>;
  return <span className="badge success">✓ {status}</span>;
}
