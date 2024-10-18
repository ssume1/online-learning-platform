import React from 'react';

export function Courses() {
    return (
        // This is a placeholder for the dashboard.
        // This will be removed/changed when an actual implementation will be made.
        <div className="p-6">
            <p className="text-2xl font-bold mb-4 text-center">This is a dummy dashboard!</p>
            <h1 className="text-2xl font-bold mb-4">Courses</h1>
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                <div className="bg-gray-100 p-4 rounded shadow">
                    <h3 className="text-xl font-semibold">Course Title 1</h3>
                    <p className="mt-2">Description of Course 1</p>
                </div>
                <div className="bg-gray-100 p-4 rounded shadow">
                    <h3 className="text-xl font-semibold">Course Title 2</h3>
                    <p className="mt-2">Description of Course 2</p>
                </div>
                <div className="bg-gray-100 p-4 rounded shadow">
                    <h3 className="text-xl font-semibold">Course Title 3</h3>
                    <p className="mt-2">Description of Course 3</p>
                </div>
            </div>
        </div>
    );
}
