#include <bits/stdc++.h>

using namespace std;

// Uncomment the following line if you need it
// typedef long long ll;

int main() {
	unsigned long long n, m, k,temp,sig,dir,sum=0,minsum=0;
	cin.sync_with_stdio(false);
	cin.tie(0);
	cout.tie(0);
	queue<unsigned long long> q1;
	queue<unsigned long long> q2;
	cin >> n >> m >> k;
	for (unsigned long long i = 1; i <= n; i++) {
		unsigned long long ai;
		cin >> ai;
		q1.push(ai);    
	}
	for (unsigned long long j=1;j<=m;j++)
	{
		dir=q1.front();
		sum+=dir;
		q1.pop();
		q2.push(dir);
	}
	temp=1;
	sig=1;
	minsum=sum;
	while(!q1.empty())
	{
		temp++;
		sum=sum+q1.front()-q2.front();
		q2.push(q1.front());
		q1.pop();
		q2.pop();
		if (sum<minsum)
		{
			minsum=sum;
			sig=temp;
		}
	}
	if (minsum < m*k)
	{
		cout << sig<<' '<<sig+m-1<<endl;
	}
    else{
    	cout<<0<<' '<<0<<endl;
	}
	return 0;
}
