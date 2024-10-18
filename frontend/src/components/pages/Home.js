export function Home({ ApiStatus }) {
    return (
        <div className="p-6">
            <p className="mt-2">API Response: {ApiStatus}</p>
        </div>
    );
}