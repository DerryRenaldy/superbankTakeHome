interface CardProps {
    children: React.ReactNode;
    className?: string;
}

interface CardContentProps {
    children: React.ReactNode;
    className?: string;
}

export function Card({ children, className }: CardProps) {
    return <div className={`bg-white shadow-md rounded-lg ${className}`}>{children}</div>;
}

export function CardContent({ children, className }: CardContentProps) {
    return <div className={`p-4 ${className}`}>{children}</div>;
}