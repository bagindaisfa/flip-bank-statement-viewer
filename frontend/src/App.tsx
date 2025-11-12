import { useState, useEffect } from "react";
import TransactionTable from "./components/TransactionTable";
import "./App.css";

type Txn = {
  timestamp: number;
  name: string;
  type: string;
  amount: number;
  status: string;
  description: string;
};

const API_BASE = import.meta.env.VITE_API_BASE || "http://localhost:8080";

function App() {
  const [file, setFile] = useState<File | null>(null);
  const [balance, setBalance] = useState<number | null>(null);
  const [transactions, setTransactions] = useState<Txn[]>([]);
  const [loading, setLoading] = useState(false);
  const [page, setPage] = useState(1);
  const [total, setTotal] = useState(0);
  const [sortBy, setSortBy] = useState("timestamp");
  const [order, setOrder] = useState<"asc" | "desc">("desc");
  const perPage = 10;

  const fetchData = async () => {
    try {
      setLoading(true);
      const [balanceRes, txnRes] = await Promise.all([
        fetch(`${API_BASE}/balance`),
        fetch(`${API_BASE}/issues?page=${page}&per_page=${perPage}&sort_by=${sortBy}&order=${order}`)
      ]);
      
      const balanceData = await balanceRes.json();
      const txnData = await txnRes.json();
      
      setBalance(balanceData.balance);
      setTransactions(txnData.items);
      setTotal(txnData.total);
    } catch (error) {
      console.error("Error fetching data:", error);
      setBalance(null);
      setTransactions([]);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData();
  }, [page, sortBy, order]);

  const handleUpload = async () => {
    if (!file) return alert("Pilih file CSV dulu");
    
    const formData = new FormData();
    formData.append("file", file);

    try {
      setLoading(true);
      const res = await fetch(`${API_BASE}/upload`, {
        method: "POST",
        body: formData,
      });
      
      if (!res.ok) {
        throw new Error("Upload gagal");
      }
      
      alert("Upload berhasil");
      setFile(null);
      await fetchData(); // Refresh both balance and transactions
    } catch (error) {
      alert(error instanceof Error ? error.message : "Terjadi kesalahan saat upload");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container">
      <div className="header">
        <h1>Statement Viewer</h1>
      </div>
      
      <div className="card">
        <div className="card-header">
          <h2 className="card-title">Upload Statement</h2>
        </div>
        <div className="flex gap-3 items-center">
          <input
            type="file"
            accept=".csv"
            onChange={(e) => setFile(e.target.files?.[0] || null)}
            className="p-2 border rounded"
            disabled={loading}
          />
          <button 
            onClick={handleUpload} 
            disabled={!file || loading}
            className="btn btn-primary"
          >
            {loading ? 'Memproses...' : 'Upload'}
          </button>
        </div>
      </div>
      
      <div className="card">
        <h2 className="card-title">Saldo: 
          <span className="font-semibold">
            {balance !== null ? `Rp${balance.toLocaleString('id-ID')}` : 'Loading...'}
          </span>
        </h2>
      </div>
      
      <div className="card">
        <TransactionTable 
          transactions={transactions}
          loading={loading}
          page={page}
          total={total}
          perPage={perPage}
          onPageChange={setPage}
          sortBy={sortBy}
          order={order}
          onSortChange={(newSortBy, newOrder) => {
            setSortBy(newSortBy);
            setOrder(newOrder);
          }}
        />
      </div>
    </div>
  );
}

export default App;
