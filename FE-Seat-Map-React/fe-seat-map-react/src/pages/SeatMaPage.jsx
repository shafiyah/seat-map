import React from 'react';
import useSeatMap from "./useSeatMap";

const SeatMapPage = () => {
  const {
    passenger,
    flight,
    selectedSeat,
    handleSave,
    renderSeatGrid
  } = useSeatMap();

  return (
    <div style={{ padding: 20, fontFamily: 'Arial' }}>
      <h3 style={{ borderBottom: '1px solid #eee', paddingBottom: 10 }}>Seat Selected</h3>
      {passenger && flight && selectedSeat && (
        <div style={{
          marginBottom: 20,
          border: '1px solid #ccc',
          borderRadius: '12px',
          padding: 20,
          backgroundColor: '#f9f9f9',
          maxWidth: 500,
          margin: 'auto'
        }}>

          <div style={{ display: 'flex', justifyContent: 'space-between', marginBottom: 10 }}>
            <strong>{passenger.FirstName}{passenger.LastName}</strong>
          </div>
          
          <div style={{ display: 'flex', justifyContent: 'space-between', marginBottom: 10 }}>
           <span>{flight.Origin} → {flight.Destination}</span>
            <span>{new Date(flight.DepartureTime).toLocaleDateString('en-GB', {
              day: '2-digit', month: 'short'
            })}, {new Date(flight.DepartureTime).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })} - {new Date(flight.ArrivalTime).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}, {flight.Duration}hr</span>
          </div>
          <div style={{ display: 'flex', justifyContent: 'space-between', textAlign: 'center' }}>
            <div>
              <div style={{ fontSize: 12 }}>Seat</div>
              <div style={{ fontWeight: 'bold', fontSize: 18 }}>{selectedSeat.Code||'-'}</div>
            </div>
            <div>
              <div style={{ fontSize: 12 }}>Position</div>
              <div style={{ fontWeight: 'bold', fontSize: 18 }}>{selectedSeat.Location || '-'}</div>
            </div>
            <div>
              <div style={{ fontSize: 12 }}>Price</div>
              <div style={{ fontWeight: 'bold', fontSize: 18 }}>{selectedSeat.CurrencyTotal} {selectedSeat.Total || '-'}</div>
            </div>
          </div>
        </div>
      )}

      <div style={{ textAlign: 'center', marginBottom: 20, marginTop: 20 }}>
        <span style={{ display: 'inline-block', backgroundColor: '#00cc44', width: 20, height: 20, marginRight: 5 }}></span> Available
        <span style={{ display: 'inline-block', backgroundColor: 'red', width: 20, height: 20, margin: '0 5px 0 30px' }}></span> Booked
      </div>

      <div style={{ marginBottom: 20 }}>{renderSeatGrid()}</div>

      <div style={{ display: 'flex', justifyContent: 'space-between', maxWidth: 500, margin: 'auto' }}>
        <div>
          <strong>Subtotal:</strong> {selectedSeat && selectedSeat.CurrencyTotal ? `${selectedSeat.CurrencyTotal} ${selectedSeat.Total}` : 'MYR 0'}
        </div>
        <button onClick={handleSave} disabled={!selectedSeat} style={{ padding: '10px 20px', backgroundColor: '#00aaff', color: '#fff', border: 'none', borderRadius: '6px' }}>
          Save
        </button>
      </div>
    </div>
  );
};

export default SeatMapPage;
