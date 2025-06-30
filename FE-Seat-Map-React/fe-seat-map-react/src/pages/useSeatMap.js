  
  import { useEffect, useState } from 'react';
  import axios from 'axios';

  const useSeatMap = () => {
  const [passenger, setPassenger] = useState(null);
  const [flight, setFlight] = useState(null);
  const [seats, setSeats] = useState([]);
  const [selectedSeat, setSelectedSeat] = useState({});
  

  useEffect(() => {
    axios.get('http://localhost:8080/passenger/1')
      .then(res => setPassenger(res.data))
      .catch(err => console.error(err));

    axios.get('http://localhost:8080/flights/1/seats')
      .then(res => {
        setFlight(res.data.flightDetail);
        setSeats(res.data.seatMaps);
      })
      .catch(err => console.error(err));
  }, []);

  const handleSeatClick = (seat) => {
    if (!seat.Available) return;
    setSelectedSeat(seat);
  };

  const handleSave = () => {
    if (!selectedSeat) return;

    axios.post('http://localhost:8080/seat/selection', {
      flightId: flight.ID,
      passengerId: passenger.ID,
      seatId: selectedSeat.ID,
    })
    .then(() => {
      alert('Seat selected successfully!');
      axios.get('http://localhost:8080/flights/1/seats')
      .then(res => {
        setSeats(res.data.seatMaps);
      })
      .catch(err => console.error(err));
    })
    
    .catch(err => console.error(err));
  };

  const renderSeatGrid = () => {
    const rows = Array.from(new Set(seats.map(s => s.Row))).sort((a, b) => a - b);
    const columns = ['A', 'B', 'C', '','D', 'E', 'F'];

    return (
      <table style={{ borderCollapse: 'collapse', margin: 'auto' }}>
        <thead>
          <tr>
            <th></th>
            {columns.map(col => <th key={col}>{col}</th>)}
          </tr>
        </thead>
        <tbody>
          {rows.map(row => (
            <tr key={row}>
              <td>{row}</td>
              {columns.map((col, idx) => {
                if (col === '') return <td key={idx} style={{ width: '50px' }}></td>; 
                const seat = seats.find(s => s.Row === row && s.Column === col);
                if (!seat) return <td key={idx}></td>;
                const isSelected = selectedSeat && selectedSeat.ID === seat.ID;
                const style = {
                  backgroundColor: !seat.Available ? 'red' : '#00cc44',
                  cursor: seat.Available ? 'pointer' : 'not-allowed',
                  padding: '10px',
                  width: '30px',
                  height: '30px',
                  borderRadius: '4px',
                  border: isSelected ? '4px solid yellow' :  '1px solid #ccc',
                };
                return (
                  <td
                    key={idx}
                    style={style}
                    onClick={() => handleSeatClick(seat)}
                  ></td>
                );
              })}
            </tr>
          ))}
        </tbody>
      </table>
    );
  };
  
  return {
    passenger,
    flight,
    seats,
    selectedSeat,
    handleSeatClick,
    handleSave,
    renderSeatGrid,
  };
};

export default useSeatMap;
