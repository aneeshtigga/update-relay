import React, { useState } from 'react';
import axios from 'axios';

function UpdateForm() {
  const [message, setMessage] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.post('http://localhost:5000/api/updates', { message });
      setMessage('');
      alert('Update successfully relayed!');
    } catch (error) {
      console.error('Error relaying update:', error);
      alert('Failed to relay update.');
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <textarea 
        value={message} 
        onChange={(e) => setMessage(e.target.value)} 
        placeholder="Enter your update"
      />
      <button type="submit">Submit</button>
    </form>
  );
}

export default UpdateForm;
